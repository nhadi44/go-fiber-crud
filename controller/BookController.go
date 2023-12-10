package controller

import (
	"fmt"
	"github.com/go-fiber-crud/database"
	"github.com/go-fiber-crud/helper"
	"github.com/go-fiber-crud/model/entity"
	request "github.com/go-fiber-crud/request/book"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookCreate(ctx *fiber.Ctx) error {
	book := new(request.BookRequest)
	if err := ctx.BodyParser(book); err != nil {
		return helper.ResponseHandler(ctx, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	// Validated Request
	validate := validator.New()
	errValidate := validate.Struct(book)

	if errValidate != nil {
		var validationErrors []ValidationErrorResponse
		for _, err := range errValidate.(validator.ValidationErrors) {
			fieldName := err.Field()
			errorMessage := fmt.Sprintf("The %s field is required!", fieldName)

			validationErrors = append(validationErrors, ValidationErrorResponse{
				Field:   fieldName,
				Message: errorMessage,
			})

			errorJson := map[string]interface{}{
				"error": validationErrors,
			}

			return helper.ResponseHandler(ctx, fiber.StatusBadRequest, "Validation error", errorJson)
		}
	}

	//handle file upload
	var fileNameString string

	fileName := ctx.Locals("fileName")
	if fileName == nil {
		return helper.ResponseHandler(ctx, fiber.StatusBadRequest, "Image cover is required", nil)
	} else {
		fileNameString = fmt.Sprintf("%v", fileName)
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  fileNameString,
	}

	errCreateBook := database.DB.Create(&newBook).Error

	if errCreateBook != nil {
		return helper.ResponseHandler(ctx, fiber.StatusBadRequest, "Failed create book", nil)
	}

	return helper.ResponseHandler(ctx, fiber.StatusOK, "Success", newBook)
}

func BookDelete(ctx *fiber.Ctx) error {
	bookId := ctx.Params("id")

	var book entity.Book

	errFindBook := database.DB.Where("id = ?", bookId).First(&book).Error
	if errFindBook != nil {
		return helper.ResponseHandler(ctx, fiber.StatusBadRequest, "Failed find book", nil)
	}

	// handle remove file
	errRemoveFile := helper.HandleRemoveFile(book.Cover, "./public/covers/")
	if errRemoveFile != nil {
		return helper.ResponseHandler(ctx, fiber.StatusInternalServerError, "Failed remove file", nil)
	}

	errDeleteBook := database.DB.Delete(&book).Error
	if errDeleteBook != nil {
		return helper.ResponseHandler(ctx, fiber.StatusBadRequest, "Failed delete book", nil)
	}

	return helper.ResponseHandler(ctx, fiber.StatusOK, "Success", nil)
}
