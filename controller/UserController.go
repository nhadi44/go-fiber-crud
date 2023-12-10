package controller

import (
	"fmt"
	"github.com/go-fiber-crud/database"
	"github.com/go-fiber-crud/helper"
	"github.com/go-fiber-crud/model/entity"
	request "github.com/go-fiber-crud/request/users"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"regexp"
)

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func UserControllerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Find(&users)

	helper.PanicHandler(result.Error)

	return helper.ResponseHandler(ctx, 200, "Success get all users", users)
}

func UserControllerCreate(ctx *fiber.Ctx) error {
	// get data from request
	user := new(request.UsersCreate)

	if err := ctx.BodyParser(user); err != nil {
		return helper.ResponseHandler(ctx, 400, "Bad Request", user)
	}

	// regex email validation
	regexEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !regexEmail.MatchString(user.Email) {
		return helper.ResponseHandler(ctx, 400, "Not valid email", nil)
	}

	// validation password length must be at least 8 characters
	if len(user.Password) < 8 {
		return helper.ResponseHandler(ctx, 400, "Password must be at least 8 characters", nil)
	}

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		var validationErrors []ValidationErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			errorMessage := fmt.Sprintf("The %s field is required!", fieldName)

			validationErrors = append(validationErrors, ValidationErrorResponse{
				Field:   fieldName,
				Message: errorMessage,
			})

			return ctx.Status(400).JSON(fiber.Map{
				"message": "Validation error",
				"errors":  validationErrors,
			})
		}
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	hashedPassword, err := helper.HasingPassword(user.Password)
	if err != nil {
		return helper.ResponseHandler(ctx, 400, "Failed create user", nil)
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return helper.ResponseHandler(ctx, 400, "Failed create user", nil)
	}

	return helper.ResponseHandler(ctx, 200, "Success create user", newUser)
}

func UserControllerGetById(ctx *fiber.Ctx) error {
	var user entity.User

	id := ctx.Params("id")

	result := database.DB.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return helper.ResponseHandler(ctx, 400, "User not found", nil)
	}

	userResponse := request.UsersGetAll{
		Id:        int(user.ID),
		Name:      user.Name,
		Address:   user.Address,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	return helper.ResponseHandler(ctx, 200, "Success get user by id", userResponse)
}

func UserControllerUpdate(ctx *fiber.Ctx) error {
	id := new(request.UserId)

	if err := ctx.BodyParser(id); err != nil {
		return helper.ResponseHandler(ctx, 400, "Bad Request", id)
	}

	var user entity.User

	result := database.DB.Where("id = ?", id.Id).First(&user)

	if result.Error != nil {
		return helper.ResponseHandler(ctx, 400, "User not found", nil)
	}

	userRequest := new(request.UsersCreate)
	// update user
	if err := ctx.BodyParser(userRequest); err != nil {
		return helper.ResponseHandler(ctx, 400, "Bad Request", userRequest)
	}

	// regex email validation
	regexEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !regexEmail.MatchString(userRequest.Email) {
		return helper.ResponseHandler(ctx, 400, "Not valid email", nil)
	}

	validate := validator.New()
	err := validate.Struct(userRequest)

	if err != nil {
		var validationErrors []ValidationErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			errorMessage := fmt.Sprintf("The %s field is required!", fieldName)

			validationErrors = append(validationErrors, ValidationErrorResponse{
				Field:   fieldName,
				Message: errorMessage,
			})

			return ctx.Status(400).JSON(fiber.Map{
				"message": "Validation error",
				"errors":  validationErrors,
			})
		}
	}

	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone

	hashedPassword, err := helper.HasingPassword(userRequest.Password)
	if err != nil {
		return helper.ResponseHandler(ctx, 400, "Failed update user", nil)
	}

	user.Password = hashedPassword

	resultUpdate := database.DB.Save(&user)

	if resultUpdate.Error != nil {
		return helper.ResponseHandler(ctx, 400, "Failed update user", nil)
	}

	return helper.ResponseHandler(ctx, 200, "Success update user", user)

}

func UserControllerDelete(ctx *fiber.Ctx) error {
	id := new(request.UserId)

	if err := ctx.BodyParser(id); err != nil {
		return helper.ResponseHandler(ctx, 400, "Bad Request", id)
	}

	var user entity.User

	result := database.DB.Where("id = ?", id.Id).First(&user)

	if result.Error != nil {
		return helper.ResponseHandler(ctx, 400, "User not found", nil)
	}

	resultDelete := database.DB.Delete(&user)

	if resultDelete.Error != nil {
		return helper.ResponseHandler(ctx, 400, "Failed delete user", nil)
	}

	return helper.ResponseHandler(ctx, 200, "Success delete user", nil)
}
