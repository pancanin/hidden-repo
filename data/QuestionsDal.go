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

func (dal QuestionsDal) Create(questionIn *models.QuestionIn) (*models.Question, error) {
	question := models.Question{
		Body:    questionIn.Body,
		Options: models.ToDal(questionIn.Options),
	}

	if err := dal.db.Create(&question).Error; err != nil {
		return nil, err
	}

	return dal.GetOne(question.ID)
}

func (dal QuestionsDal) GetOne(questionId uint) (*models.Question, error) {
	question := models.Question{}

	if err := dal.db.Preload("Options").First(&question, questionId).Error; err != nil {
		return nil, err
	}

	return &question, nil
}

func (dal QuestionsDal) GetAll() ([]models.Question, error) {
	var questions []models.Question

	if err := dal.db.Model(models.Question{}).Preload("Options").Find(&questions).Error; err != nil {
		return questions, err
	}

	return questions, nil
}

func (dal QuestionsDal) Update(questionId uint, question *models.QuestionIn) error {
	return dal.db.Transaction(func(tx *gorm.DB) error {
		err := tx.
			Model(models.Question{}).
			Where("id = ?", questionId).
			Updates(question.ToDal()).
			Error

		if err != nil {
			return err
		}

		if err := tx.Where("question_id = ?", questionId).Delete(models.Option{}).Error; err != nil {
			return err
		}

		dalOptions := models.ToDal(question.Options)

		for idx := range dalOptions {
			dalOptions[idx].QuestionID = questionId
		}

		if err := tx.Create(dalOptions).Error; err != nil {
			return err
		}

		return nil
	})
}
