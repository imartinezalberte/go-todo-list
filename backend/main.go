package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/imartinezalberte/go-todo-list/backend/internal/tag"
	"github.com/imartinezalberte/go-todo-list/backend/utils"
)

// TODO: This can't be here...
const URI string = "host=localhost user=ivan password=abc123. dbname=todo port=5432 sslmode=disable TimeZone=Europe/Madrid"

func main() {
	db := utils.GetOrPanic(gorm.Open(postgres.Open(URI), &gorm.Config{}))

	if err := db.AutoMigrate(new(tag.Tag)); err != nil {
		panic(err)
	}

	r := gin.Default()

	tag.Handler(tag.NewService(tag.NewRepo(db))).Routes(r)

	r.Run(":8080")
}
