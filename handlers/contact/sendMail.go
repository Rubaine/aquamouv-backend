package contact

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendMailAsync(destinationEmail string, firstName string, lastName string) {
	go func() {
		auth := smtp.PlainAuth(
			"",
			os.Getenv("MAIL_USER"),
			os.Getenv("MAIL_PASSWORD"),
			os.Getenv("MAIL_HOST"),
		)

		msg := fmt.Sprintf(`Subject: Contact Information Aquamouv

Bonjour %s %s,

Nous avons bien reçu votre demande pour une séance découverte gratuite. Le centre Aquamouv de Senlis est heureux de vous accueillir pour 
vous faire essayer un cours d'aquabiking coaché par des professeurs de sport diplômés.

Prenez vite rendez-vous avec le centre Aquamouv de Senlis en appelant au 06 20 52 17 37. Pour effectuer votre séance, pensez à prendre un 
maillot de bain, une serviette de bain, un gel douche, une bouteille d'eau et des chaussures d'aquavélo.

Horaires d'ouverture: Du Lundi au vendredi 9h-20h / Samedi 9h-16h / Dimanche 9h-15h / Jours Fériés horaires du week-end 
Adresse: 6H Av. du Poteau, 60300 Chamant, France (à côté du Norauto)

*Offre non cumulable.

Cordialement,
L'équipe Aquamouv`, lastName, firstName)

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
