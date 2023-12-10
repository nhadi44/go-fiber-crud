package helper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathImage = "./public/covers/"

func HandleSingleFile(ctx *fiber.Ctx) error {
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error upload file", errFile.Error())
	}

	var fileName string
	var newFilename string

	if file != nil {
		contentTypeFile := file.Header.Get("Content-Type")
		if contentTypeFile != "image/png" && contentTypeFile != "image/jpeg" {
			return ResponseHandler(ctx, fiber.StatusBadRequest, "File must be image", nil)
		} else {
			fileName = file.Filename
			extensionFile := filepath.Ext(fileName)
			newFilename = GenerateRandomString(10) + extensionFile

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", newFilename))
			if errSaveFile != nil {
				log.Println("Error save file", errSaveFile.Error())
			}
		}

	} else {
		log.Println("File not found")
	}

	if fileName != "" {
		ctx.Locals("fileName", newFilename)
	} else {
		ctx.Locals("fileName", nil)
	}

	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errorForm := ctx.MultipartForm()

	if errorForm != nil {
		log.Println("Error Read Multiple File", errorForm.Error())
	}

	files := form.File["photos"]

	var fileNames []string

	for i, file := range files {
		var fileName string
		if file != nil {
			extensionFile := filepath.Ext(file.Filename)
			newFileName := GenerateRandomString(10) + extensionFile
			fileName = fmt.Sprintf("%d%s", i, newFileName)

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", fileName))
			if errSaveFile != nil {
				log.Println("Error save file", errSaveFile.Error())
			}
		} else {
			log.Println("File not found")
		}

		if fileName != "" {
			fileNames = append(fileNames, fileName)
			ctx.Locals("fileName", fileName)
		} else {
			ctx.Locals("fileName", nil)
		}
	}

	ctx.Locals("fileNames", fileNames)

	return ctx.Next()
}

func HandleRemoveFile(filename string, pathFile ...string) error {

	if len(pathFile) > 0 {
		err := os.Remove(pathFile[0] + filename)

		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	} else {
		err := os.Remove(DefaultPathImage + filename)

		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	}

	return nil
}
