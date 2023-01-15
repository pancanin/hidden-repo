package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("API_PORT")

	if len(port) == 0 {
		port = "3000"
	}

	portParam := fmt.Sprintf(":%s", port)

	fmt.Printf("Server listening on port %s", portParam)

	http.ListenAndServe(portParam, r)
}
