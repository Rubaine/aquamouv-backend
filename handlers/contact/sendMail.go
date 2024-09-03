package contact

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

func sendMailAsync(destinationEmail string, firstName string, lastName string) {
	go func() {
		auth := smtp.PlainAuth(
			"",
			os.Getenv("MAIL_USER"),
			os.Getenv("MAIL_PASSWORD"),
			os.Getenv("MAIL_HOST"),
		)

		msg := fmt.Sprintf(`From: %s
To: %s
Subject: =?UTF-8?B?%s?=
MIME-Version: 1.0
Content-Type: text/plain; charset="utf-8"
Content-Transfer-Encoding: 8bit

Bonjour %s %s,

Nous avons bien reçu votre demande pour une séance découverte gratuite. Le centre Aquamouv de Senlis est heureux de vous accueillir pour 
vous faire essayer un cours d'aquabiking coaché par des professeurs de sport diplômés.

Prenez vite rendez-vous avec le centre Aquamouv de Senlis en appelant au 06 20 52 17 37. Pour effectuer votre séance, pensez à prendre un 
maillot de bain, une serviette de bain, un gel douche, une bouteille d'eau et des chaussures d'aquavélo.

Horaires d'ouverture: Du Lundi au vendredi 9h-20h / Samedi 9h-16h / Dimanche 9h-15h / Jours Fériés horaires du week-end 
Adresse: 6H Av. du Poteau, 60300 Chamant, France (à côté du Norauto)

*Offre non cumulable.

Cordialement,
L'équipe Aquamouv`, os.Getenv("MAIL_USER"), destinationEmail, encodeRFC2047("Contact Information Aquamouv"), lastName, firstName)

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

func sendMailToManager(destinationEmail string, firstName string, lastName string, phoneNumber string, email string) {
	go func() {
		auth := smtp.PlainAuth(
			"",
			os.Getenv("MAIL_USER"),
			os.Getenv("MAIL_PASSWORD"),
			os.Getenv("MAIL_HOST"),
		)

		msg := fmt.Sprintf(`From: %s
To: %s
Subject: =?UTF-8?B?%s?=
MIME-Version: 1.0
Content-Type: text/plain; charset="utf-8"
Content-Transfer-Encoding: 8bit

Bonjour Patricia,

M. %s %s a demandé à faire une séance d'essai. Merci de le rappeler au plus vite au: %s ou de le contacter par mail à: %s.

Cordialement,
L'équipe Aquamouv`, os.Getenv("MAIL_USER"), destinationEmail, encodeRFC2047("Nouvelle demande de séance d'essai"), lastName, firstName, phoneNumber, email)

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

// encodeRFC2047 encode un string pour être compatible avec le format de sujet d'email
func encodeRFC2047(String string) string {
	// UTF-8 avec base64 encoding
	addr := mail.Address{Name: String, Address: ""}
	return strings.Trim(addr.String(), " <>")
}
