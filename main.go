package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	dals "questions/data"
	models "questions/data/models"

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

	option1 := models.OptionUpdate{
		ID:      1,
		Body:    "360 MT",
		Correct: false,
	}
	options2 := models.OptionUpdate{
		ID:      2,
		Body:    "4500 MT",
		Correct: true,
	}
	options3 := models.OptionUpdate{
		Body:    "1000 MT",
		Correct: true,
	}
	question := models.QuestionUpdate{
		Body:    "What is the weight of the jupiter planet?",
		Options: []models.OptionUpdate{option1, options2, options3},
	}

	questionsDal.Update(1, &question)

	// questions, errs := questionsDal.GetAll()

	// if errs != nil {
	// 	log.Fatal("Ofatali se neshto!")
	// }

	// for _, quest := range questions {
	// 	fmt.Println(quest.Body)

	// 	for _, opt := range quest.Options {
	// 		fmt.Println("- " + opt.Body)
	// 	}
	// }

	r := gin.Default()

	// r.POST("/todo", todoHandler.Create)
	// r.GET("/todo", todoHandler.GetAll)
	// r.GET("/todo/:id", todoHandler.Get)
	// r.PUT("/todo", todoHandler.Update)
	// r.DELETE("/todo/:id", todoHandler.Delete)

	http.ListenAndServe(":3000", r)
}
