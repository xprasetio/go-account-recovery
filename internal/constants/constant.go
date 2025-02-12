package constants

const (
	RecoveryEmailTemplate = `
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body {
					font-family: Arial, sans-serif;
					line-height: 1.6;
					color: #333333;
				}
				.container {
					max-width: 600px;
					margin: 0 auto;
					padding: 20px;
				}
				.header {
					background-color:rgb(144, 185, 226);
					padding: 20px;
					text-align: center;
					border-radius: 5px;
				}
				.content {
					padding: 20px;
				}
				.code {
					font-size: 24px;
					font-weight: bold;
					text-align: center;
					padding: 10px;
					margin: 20px 0;
					background-color:rgb(80, 92, 104);
					border-radius: 5px;
				}
				.footer {
					text-align: center;
					font-size: 12px;
					color: #666666;
					margin-top: 20px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h2>Account Recovery Code</h2>
				</div>
				<div class="content">
					<p>Hello,</p>
					<p>We received a request to recover your account. Here is your recovery code:</p>
					<div class="code">{{.Code}}</div>
					<p>Please use this code to reset your password.</p>
					<p>For security reasons, do not share this code with anyone.</p>
				</div>
				<div class="footer">
					<p>This is an automated message, please do not reply to this email.</p>
					<p>&copy; {{.Year}} PT Digital Security Indonesia. All rights reserved.</p>
				</div>
			</div>
		</body>
		</html>
		`
	SuccessMessage      = "success"
	ErrFailedBadRequest = "data tidak sesuai"
	ErrServerError      = "terjadi kesalahan pada server"
)