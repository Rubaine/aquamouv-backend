package auth

import (
	"remy-aquavelo/config"
	"remy-aquavelo/models"

	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(ctx iris.Context){
	var user models.User

	if err := ctx.ReadJSON(&user); err != nil{
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "request body is not a valid JSON"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "failed to hash the password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := config.Cfg.DB.Create(&user).Error; err != nil{
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "failed to create a user"})
		return
	}

	ctx.StatusCode(iris.StatusOK)
    ctx.JSON(iris.Map{"message": "User registered successfully", "user": user})
}