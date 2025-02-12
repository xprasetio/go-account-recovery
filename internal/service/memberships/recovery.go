package memberships

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"text/template"
	"time"

	"xprasetio/go-account-recovery.git/internal/constants"
	"xprasetio/go-account-recovery.git/internal/models/memberships"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	mail "github.com/xhit/go-simple-mail/v2"
)
 
func (s *service) InitiateRecovery(ctx context.Context, request memberships.ResetEmailRequest) error {
    userDetail, err := s.repository.GetUser(ctx, request.Email, "", uint(0),"")
    if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Error get user from database")
		return  err
	}
	if userDetail == nil {
		return  errors.New("email not found")
	}
    // Generate recovery code
    code := make([]byte, 32/2)
    if _, err := rand.Read(code); err != nil {
        return err
    }
    recoveryCode := hex.EncodeToString(code)
    userDetail.RecoverCode = recoveryCode
    if err := s.repository.UpdateUser(ctx,userDetail); err != nil {
        return err
    }
    // Send email
    return s.sendRecoveryEmail(request.Email, recoveryCode)
}

// TemplateData holds data for email template
type TemplateData struct {
    Code string
    Year string
}
func (s *service) sendRecoveryEmail(email, code string) error {
    // Create new email client
    server := mail.NewSMTPClient()

    cfg := s.cfg
    // Set server settings
    server.Host = cfg.EmailConfig.SMTPHost //emailConfig.SMTPHost
    server.Port = cfg.EmailConfig.SMTPPort
    server.Username = cfg.EmailConfig.SMTPUsername
    server.Password = cfg.EmailConfig.SMTPPassword
    server.Encryption = mail.EncryptionSTARTTLS // or EncryptionSSL based on your needs
    
    // Set connection timeout
    server.ConnectTimeout = 10 * time.Second
    server.SendTimeout = 10 * time.Second
    
    // Create SMTP client
    smtpClient, err := server.Connect()
    if err != nil {
        return fmt.Errorf("failed to connect to SMTP server: %v", err)
    }
    // Create email
    email_msg := mail.NewMSG()
    email_msg.SetFrom(fmt.Sprintf("%s <%s>", cfg.EmailConfig.FromName, cfg.EmailConfig.FromEmail))
    email_msg.AddTo(email)
    email_msg.SetSubject("Account Recovery Code")

    // Parse email template
    tmpl, err := template.New("recovery_email").Parse(constants.RecoveryEmailTemplate)
    if err != nil {
        return fmt.Errorf("failed to parse email template: %v", err)
    }

    // Prepare template data
    data := TemplateData{
        Code: code,
        Year: time.Now().Format("2006"), // Dynamic year
    }

    // Execute template with data
    var body bytes.Buffer
    if err := tmpl.Execute(&body, data); err != nil {
        return fmt.Errorf("failed to execute email template: %v", err)
    }

    // Set email body
    email_msg.SetBody(mail.TextHTML, body.String())

    // Send email
    if err := email_msg.Send(smtpClient); err != nil {
        return fmt.Errorf("failed to send email: %v", err)
    }

    // Log success (you might want to use proper logging in production)
    fmt.Printf("Recovery code sent successfully to %s\n", email)
    return nil
}