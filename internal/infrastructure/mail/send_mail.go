package mail

import (
	"crypto/tls"
	"fmt"
	"log" // Import the log package
	"os"

	"github.com/danubiobwm/goEmailN/internal/domain/campaign"
	"gopkg.in/gomail.v2"
)

func SendMail(campaign *campaign.Campaign) error {
	fmt.Println("Sending mail...")

	// Debug: Print environment variables (remove in production!)
	// fmt.Println("EMAIL_SMTP:", os.Getenv("EMAIL_SMTP"))
	// fmt.Println("EMAIL_USER:", os.Getenv("EMAIL_USER"))
	// fmt.Println("EMAIL_PASSWORD:", os.Getenv("EMAIL_PASSWORD")) // Be very careful with logging passwords

	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	// Important: Force TLS connection (Gmail often requires it)
	d.TLSConfig = &tls.Config{ServerName: "smtp.gmail.com"} // Import "crypto/tls"

	var emails []string
	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", campaign.Name)
	m.SetBody("text/html", campaign.Content)

	if err := d.DialAndSend(m); err != nil {
		// CRITICAL: Log the actual error!
		fmt.Println("Error sending email:", err) // For debugging
		log.Println("Error sending email:", err) // Use log package for more robust logging
		return err                               // Or handle the error as needed
	}

	fmt.Println("Email sent successfully!") // Confirmation message
	return nil
}
