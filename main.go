package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	auth "questions/auth"
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

	dals.NewUsersDal(db)
	questionsDal := dals.NewQuestionsDal(db)
	questionHandler := handlers.NewQuestionHandler(&questionsDal)

	r := gin.Default()

	r.Use(auth.JWTTokenAuthMiddleware())

	r.POST("/question", questionHandler.Create)
	r.GET("/questions", questionHandler.GetPaginated)
	r.PUT("/question/:id", questionHandler.Update)
	r.DELETE("/question/:id", questionHandler.Delete)

	port := os.Getenv("API_PORT")

	if len(port) == 0 {
		port = "3000"
	}

	portParam := fmt.Sprintf(":%s", port)

	fmt.Printf("Server listening on port %s\n", portParam)

	http.ListenAndServe(portParam, r)
}
