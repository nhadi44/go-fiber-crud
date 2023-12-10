package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-fiber-crud/database"
	"github.com/go-fiber-crud/helper"
	"github.com/go-fiber-crud/model/entity"
	request "github.com/go-fiber-crud/request/auth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func LoginController(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return helper.ResponseHandler(ctx, 400, "Bad Request", loginRequest)
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)

	if errValidate != nil {
		return helper.ResponseHandler(ctx, 400, errValidate.Error(), loginRequest)
	}

	// check available user
	var user entity.User
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return helper.UnAuthorizedResponse(ctx, "Email or Password is wrong")
	}

	// check validation password
	isValid := helper.CheckPassword(loginRequest.Password, user.Password)
	if !isValid {
		return helper.UnAuthorizedResponse(ctx, "Email or Password is wrong")
	}

	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	if user.Email == "hadi1@mail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errorGenerateToken := helper.GenerateJwtToken(&claims)
	if errorGenerateToken != nil {
		log.Println(errorGenerateToken)
		return helper.UnAuthorizedResponse(ctx, "Token generation failed")
	}

	newJson := map[string]interface{}{
		"token": token,
	}

	return helper.ResponseHandler(ctx, 200, "Login success", newJson)
}
