package memberships

import "gorm.io/gorm"

type (
	User struct { 
		gorm.Model
		Email     string `gorm:"unique;not null"`
		Username  string `gorm:"unique;not null"`
		Password  string `gorm:"not null"`
		RecoverCode string `gorm:"null"`
		CreatedBy string `gorm:"not null"`
		UpdatedBy string `gorm:"not null"`
	}
)

type (
	SignUpRequest struct {		
		Email       string `json:"email"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		RecoverCode string `json:"recover_code" gorm:"default:coderecover"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	
	ResetEmailRequest struct {
		Email string `json:"email"`
	}

	ResetCodeRequest struct {
		Code string `json:"recover_code"`
	}
	ResetPasswordRequest struct {
		Password string `json:"password"`
	}
)

type (
	LoginResponse struct {	
		AccessToken string `json:"accessToken"`
	}
)