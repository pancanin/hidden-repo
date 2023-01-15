package data

import (
	models "questions/data/models"

	"gorm.io/gorm"
)

const (
	PAGING_MAX_PAGES         = 100
	PAGING_DEFAULT_PAGE_SIZE = 10
)

type QuestionsDal struct {
	db *gorm.DB
}

func NewQuestionsDal(db *gorm.DB) QuestionsDal {
	db.AutoMigrate(&models.Question{}, &models.Option{})
	db.Exec("PRAGMA foreign_keys = ON;")

	return QuestionsDal{db}
}

func (dal QuestionsDal) Create(params models.QuestionCreateParams) (*models.Question, error) {
	question := models.Question{
		Body:    params.Question.Body,
		Options: models.ToDal(params.Question.Options),
		UserID:  params.UserID,
	}

	if err := dal.db.Create(&question).Error; err != nil {
		return nil, err
	}

	return dal.GetOne(models.QuestionGetOneParams{
		QuestionID: question.ID,
		UserID:     params.UserID,
	})
}

func (dal QuestionsDal) GetOne(params models.QuestionGetOneParams) (*models.Question, error) {
	question := models.Question{}

	if err := dal.db.Preload("Options").
		Where("id = ? AND user_id = ?", params.QuestionID, params.UserID).
		First(&question).
		Error; err != nil {

		return nil, err
	}

	return &question, nil
}

func (dal QuestionsDal) GetPaginated(params models.QuestionsGetPaginatedParams) ([]models.Question, error) {
	var questions []models.Question

	if err := dal.db.
		Scopes(dal.Paginate(params.Req, PAGING_MAX_PAGES, PAGING_DEFAULT_PAGE_SIZE)).
		Model(models.Question{}).
		Preload("Options").
		Where("user_id = ?", params.UserID).
		Find(&questions).Error; err != nil {

		return questions, err
	}

	return questions, nil
}

func (dal QuestionsDal) Update(params models.QuestionUpdateParams) error {
	return dal.db.Transaction(func(tx *gorm.DB) error {
		err := tx.
			Model(models.Question{}).
			Where("id = ? AND user_id = ?", params.QuestionID, params.UserID).
			Updates(params.Question.ToDal()).
			Error

		if err != nil {
			return err
		}

		if err := tx.Where("question_id = ?", params.QuestionID).Delete(models.Option{}).Error; err != nil {
			return err
		}

		dalOptions := models.ToDal(params.Question.Options)

		for idx := range dalOptions {
			dalOptions[idx].QuestionID = params.QuestionID
		}

		if err := tx.Create(dalOptions).Error; err != nil {
			return err
		}

		return nil
	})
}

func (dal QuestionsDal) Delete(params models.QuestionDeleteParams) error {
	return dal.db.
		Where("id = ? AND user_id = ?", params.QuestionID, params.UserID).
		Delete(models.Question{}).Error
}
