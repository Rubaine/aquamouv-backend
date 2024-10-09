package contact

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/kataras/golog"
)

func sendMailAsync(destinationEmail string, firstName string, lastName string) {
	go func() {
		golog.Info("Starting sendMailAsync goroutine")

		auth := smtp.PlainAuth(
			"",
			os.Getenv("MAIL_USER"),
			os.Getenv("MAIL_PASSWORD"),
			os.Getenv("MAIL_HOST"),
		)

		subject := "Contact Information Aquamouv"
		body := fmt.Sprintf(`Bonjour %s %s,

Nous avons bien reçu votre demande pour une séance découverte gratuite. Le centre Aquamouv de Senlis est heureux de vous accueillir pour 
vous faire essayer un cours d'aquabiking coaché par des professeurs de sport diplômés.

Prenez vite rendez-vous avec le centre Aquamouv de Senlis en appelant au 06 20 52 17 37. Pour effectuer votre séance, pensez à prendre un 
maillot de bain, une serviette de bain, un gel douche, une bouteille d'eau et des chaussures d'aquavélo.

Horaires d'ouverture: Du Lundi au vendredi 9h-20h / Samedi 9h-16h / Dimanche 9h-15h / Jours Fériés horaires du week-end 
Adresse: 6H Av. du Poteau, 60300 Chamant, France (à côté du Norauto)

*Offre non cumulable.

Cordialement,
L'équipe Aquamouv`, lastName, firstName)

		msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"utf-8\"\r\n\r\n%s",
			os.Getenv("MAIL_USER"), destinationEmail, subject, body)

		golog.Info("Sending email to:", destinationEmail)
		err := smtp.SendMail(
			os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"),
			auth,
			os.Getenv("MAIL_USER"),
			[]string{destinationEmail},
			[]byte(msg),
		)

		if err != nil {
			golog.Error("Failed to send email:", err)
		} else {
			golog.Info("Email sent successfully to:", destinationEmail)
		}
	}()
}

func sendMailToManager(destinationEmail string, firstName string, lastName string, phoneNumber string, email string) {
	go func() {
		golog.Info("Starting sendMailToManager goroutine")

		auth := smtp.PlainAuth(
			"",
			os.Getenv("MAIL_USER"),
			os.Getenv("MAIL_PASSWORD"),
			os.Getenv("MAIL_HOST"),
		)

		subject := "Nouvelle demande de séance d'essai"
		body := fmt.Sprintf(`Bonjour Patricia,

M. %s %s a demandé à faire une séance d'essai. Merci de le rappeler au plus vite au: %s ou de lui envoyer un email à: %s.

Cordialement,
L'équipe Aquamouv`, lastName, firstName, phoneNumber, email)

		msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"utf-8\"\r\n\r\n%s",
			os.Getenv("MAIL_USER"), destinationEmail, subject, body)

		golog.Info("Sending email to manager:", destinationEmail)
		err := smtp.SendMail(
			os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"),
			auth,
			os.Getenv("MAIL_USER"),
			[]string{destinationEmail},
			[]byte(msg),
		)

		if err != nil {
			golog.Error("Failed to send email to manager:", err)
		} else {
			golog.Info("Email sent successfully to manager:", destinationEmail)
		}
	}()
}
