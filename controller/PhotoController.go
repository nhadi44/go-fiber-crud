package controller

import (
	"fmt"
	"log"

	"github.com/go-fiber-crud/database"
	"github.com/go-fiber-crud/helper"
	"github.com/go-fiber-crud/model/entity"
	request "github.com/go-fiber-crud/request/photo"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PhotoControllerCreate(ctx *fiber.Ctx) error {
	photo := new(request.PhotoRequest)
	if err := ctx.BodyParser(photo); err != nil {
		return helper.ResponseHandler(ctx, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	// Validated Request
	validate := validator.New()
	errValidate := validate.Struct(photo)

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
	// var fileNameString string

	fileNames := ctx.Locals("fileNames")
	if fileNames == nil {
		return helper.ResponseHandler(ctx, fiber.StatusBadRequest, "Image cover is required", nil)
	} else {
		// fileNameString = fmt.Sprintf("%v", fileNames)
		fileNamesData := fileNames.([]string)

		for _, fileName := range fileNamesData {
			newPhoto := entity.Photo{
				Image:      fileName,
				CategoryId: photo.CategoryId,
			}

			errCreateBook := database.DB.Create(&newPhoto).Error
			if errCreateBook != nil {
				return helper.ResponseHandler(ctx, fiber.StatusInternalServerError, "Failed create photo", errCreateBook.Error())
			}
		}
	}

	// log.Println(fileNameString)
	return helper.ResponseHandler(ctx, fiber.StatusOK, "Success", nil)
}

func PhotoControllerDelete(ctx *fiber.Ctx) error {
	photoId := ctx.Params("id")

	var photo entity.Photo

	err := database.DB.Debug().First(&photo, "id = ?", photoId).Error
	if err != nil {
		return helper.ResponseHandler(ctx, fiber.StatusBadRequest, "Photo not found", err.Error())
	}

	// handle remove file
	errDeleteFile := helper.HandleRemoveFile(photo.Image)
	if errDeleteFile != nil {
		log.Println("Failed to remove file")
	}

	errDelete := database.DB.Debug().Delete(&photo).Error
	if errDelete != nil {
		return helper.ResponseHandler(ctx, fiber.StatusInternalServerError, "Failed delete photo", errDelete.Error())
	}

	return helper.ResponseHandler(ctx, fiber.StatusOK, "Success", nil)
}
