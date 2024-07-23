package contact

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendMailAsync(destinationEmail string) {
	go func() {
		auth := smtp.PlainAuth(
			"",
			os.Getenv("MAIL_USER"),
			os.Getenv("MAIL_PASSWORD"),
			os.Getenv("MAIL_HOST"),
		)

		msg := "Subject: Contact Information\n\nThank you for contacting us. We will get back to you as soon as possible."
		err := smtp.SendMail(
			os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"),
			auth,
			os.Getenv("MAIL_USER"),
			[]string{destinationEmail},
			[]byte(msg),
		)

		if err != nil {
			fmt.Println(err)
		}
	}()
}