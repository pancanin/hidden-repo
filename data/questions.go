package data

import (
	models "questions/data/models"

	"gorm.io/gorm"
)

type QuestionsDal struct {
	db *gorm.DB
}

func NewQuestionsDal(db *gorm.DB) QuestionsDal {
	db.AutoMigrate(&models.Question{}, &models.Option{})
	return QuestionsDal{db}
}

func (dal QuestionsDal) Create(questionIn *models.QuestionIn) (models.Question, error) {
	todo := models.Question{
		Body:    questionIn.Body,
		Options: models.ToDal(questionIn.Options),
	}

	if err := dal.db.Create(&todo).Error; err != nil {
		return models.Question{}, err
	}

	return todo, nil
}

func (dal QuestionsDal) GetAll() ([]models.Question, error) {
	var questions []models.Question

	if err := dal.db.Model(models.Question{}).Preload("Options").Find(&questions).Error; err != nil {
		return questions, err
	}

	return questions, nil
}

func (dal QuestionsDal) Update(questionId uint, question *models.QuestionUpdate) error {
	return dal.db.Transaction(func(tx *gorm.DB) error {
		err := dal.db.
			Model(models.Question{}).
			Where("id = ?", questionId).
			Updates(question.ToDal()).
			Error

		if err != nil {
			return err
		}

		for _, option := range question.Options {
			dalOption := option.ToDal()

			// Create a method for that
			if dal.db.Model(models.Option{}).Where("id = ?", option.ID).Updates(&dalOption).RowsAffected == 0 {
				dal.db.Create(&dalOption)
			}
		}
		return nil
	})
}

// func (dal TodoDal) Delete(id int) error {
// 	return dal.db.
// 		Delete(&models.Question{}, id).
// 		Error
// }
