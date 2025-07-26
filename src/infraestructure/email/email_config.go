package email

import "os"

type EmailConfig struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	From     string
}

func GetEmailConfig() EmailConfig {
	return EmailConfig{
		SMTPHost: os.Getenv("EMAIL_SMTP"),
		SMTPPort: os.Getenv("EMAIL_PORT"),
		Username: os.Getenv("EMAIL_USERNAME"),
		Password: os.Getenv("EMAIL_PASSWORD"),
		From:     os.Getenv("EMAIL_FROM"),
	}
}
