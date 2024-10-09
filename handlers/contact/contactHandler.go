package contact

import (
	"os"
	"remy-aquavelo/config"
	"remy-aquavelo/models"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

func ContactSubmitHandler(ctx iris.Context) {
	var user models.ContactInfo

	golog.Info("Received a contact submission request")

	if err := ctx.ReadJSON(&user); err != nil {
		golog.Error("Failed to read JSON from request body: ", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "request body is not a valid JSON"})
		return
	}

	golog.Info("Contact information read successfully: ", user)

	if err := config.Cfg.DB.Create(&user).Error; err != nil {
		golog.Error("Failed to store contact information: ", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "failed to store contact information"})
		return
	}

	golog.Info("Contact information stored successfully in the database")

	sendMailAsync(user.Email, user.FirstName, user.LastName)
	golog.Info("Sent confirmation email to user: ", user.Email)

	sendMailToManager(os.Getenv("MAIL_USER"), user.FirstName, user.LastName, user.Phone, user.Email)
	golog.Info("Sent notification email to manager")

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Contact information stored successfully"})
	golog.Info("Response sent to client")
}
