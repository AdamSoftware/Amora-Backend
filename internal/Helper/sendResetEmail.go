package helper

import (
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)


func SendResetEmail(toEmail, resetLink string) error {
   
	SendGridAPIKey := os.Getenv("SENDGRID_API_KEY")

	if SendGridAPIKey == "" {
		return errors.New("SENDGRID_API_KEY is not set in environment variables")
	}

	from := mail.NewEmail("Amora Support", "no-reply@yourapp.com")
	subject := "Password Reset Request"
	to := mail.NewEmail("User", toEmail)
	plainTextContent := "You requested a password reset. Click the link below to reset your password:\n" + resetLink
	htmlContent := "<strong>You requested a password reset. Click the link below to reset your password:</strong><br><a href=\"" + resetLink + "\">Reset Password</a>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(SendGridAPIKey)
	_, err := client.Send(message)
	return err
}

