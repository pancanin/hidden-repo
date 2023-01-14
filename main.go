package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	dals "questions/data"
	handlers "questions/http-handlers"
)

func main() {
	db, err := gorm.Open(sqlite.Open("questions.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to DB")
	}

	todoDal := dals.NewTodoDal(db)
	todoHandler := handlers.NewTodoHandler(&todoDal)

	r := gin.Default()

	r.POST("/todo", todoHandler.Create)
	r.GET("/todo", todoHandler.GetAll)
	r.GET("/todo/:id", todoHandler.Get)
	r.PUT("/todo", todoHandler.Update)
	r.DELETE("/todo/:id", todoHandler.Delete)

	http.ListenAndServe(":3000", r)
}
