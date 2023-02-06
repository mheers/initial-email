package main

import (
	"log"
	"net/smtp"
	"os"
)

func main() {
	// Read environment variables
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("SMTP_FROM")
	fromName := os.Getenv("SMTP_FROM_NAME")
	to := os.Getenv("SMTP_TO")
	toName := os.Getenv("SMTP_TO_NAME")
	subject := os.Getenv("SMTP_SUBJECT")
	text := os.Getenv("SMTP_TEXT")

	log.Println("Sending email...")
	log.Println("SMTP_HOST:", smtpHost)
	log.Println("SMTP_PORT:", smtpPort)
	log.Println("SMTP_USERNAME:", username)
	log.Println("SMTP_FROM:", from)
	log.Println("SMTP_FROM_NAME:", fromName)
	log.Println("SMTP_TO:", to)
	log.Println("SMTP_TO_NAME:", toName)
	log.Println("SMTP_SUBJECT:", subject)
	log.Println("SMTP_TEXT:", text)

	// Build the email
	auth := smtp.PlainAuth("", username, password, smtpHost)
	message := []byte("From: " + fromName + " <" + from + ">\n" +
		"To: " + toName + " <" + to + ">\n" +
		"Subject: " + subject + "\n\n" +
		text)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		log.Fatalf("Failed to send email: %s", err)
	}

	log.Println("Email sent successfully.")
}
