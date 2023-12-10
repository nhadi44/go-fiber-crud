package migration

import (
	"fmt"
	"github.com/go-fiber-crud/database"
	"github.com/go-fiber-crud/helper"
	"github.com/go-fiber-crud/model/entity"
)

func MigrationInit() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Category{}, &entity.Photo{})

	helper.PanicHandler(err)

	fmt.Println("Migration completed")
}
