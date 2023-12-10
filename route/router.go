package route

import (
	"github.com/go-fiber-crud/config"
	"github.com/go-fiber-crud/controller"
	"github.com/go-fiber-crud/helper"
	"github.com/go-fiber-crud/middleware"
	"github.com/gofiber/fiber/v2"
)

func Router(router *fiber.App) {
	router.Get("/user", middleware.Auth, controller.UserControllerGetAll)
	router.Get("/user/:id", controller.UserControllerGetById)
	router.Post("/user", controller.UserControllerCreate)
	router.Put("/user", controller.UserControllerUpdate)
	router.Delete("/user", controller.UserControllerDelete)

	router.Static("/public", config.ProjectRootPath+"/public/assets")

	router.Post("/login", controller.LoginController)

	router.Post("/book", helper.HandleSingleFile, controller.BookCreate)
	router.Delete("/book/:id", controller.BookDelete)
	router.Post("/gallery", helper.HandleMultipleFile, controller.PhotoControllerCreate)
	router.Delete("/gallery/:id", controller.PhotoControllerDelete)
}
