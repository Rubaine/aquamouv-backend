package contact

import (
	"remy-aquavelo/config"
	"remy-aquavelo/models"

	"github.com/kataras/iris/v12"
)

func ContactSubmitHandler(ctx iris.Context) {
	var user models.ContactInfo

	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "request body is not a valid JSON"})
		return
	}

	if err := config.Cfg.DB.Create(&user).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "failed to store contact information"})
		return
	}

	sendMailAsync(user.Email, user.FirstName, user.LastName)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Contact information stored successfully"})
}
