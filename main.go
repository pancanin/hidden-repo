package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	dals "questions/data"
	handlers "questions/http-handlers"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("questions.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to DB")
	}

	questionsDal := dals.NewQuestionsDal(db)
	//todoHandler := handlers.NewTodoHandler(&questionsDal)
	questionHandler := handlers.NewQuestionHandler(&questionsDal)

	r := gin.Default()

	r.POST("/question", questionHandler.Create)
	r.GET("/questions", questionHandler.GetAll)
	r.PUT("/question/:id", questionHandler.Update)
	// r.DELETE("/todo/:id", todoHandler.Delete)

	http.ListenAndServe(":3000", r)
}
