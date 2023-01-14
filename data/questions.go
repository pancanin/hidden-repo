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

		if err := dal.handleDeletedOptions(questionId, question.Options); err != nil {
			return err
		}

		// Handle existing and new options
		for _, option := range question.Options {
			dalOption := option.ToDal(questionId)

			if dal.db.Model(models.Option{}).Where("id = ? AND question_id = ?", option.ID, questionId).Updates(&dalOption).RowsAffected == 0 {
				dal.db.Create(&dalOption)
			}
		}

		return nil
	})
}

func (dal *QuestionsDal) handleDeletedOptions(questionId uint, optionsUpsert []models.OptionUpsert) error {
	var options []models.Option

	if err := dal.db.Model(models.Option{}).Find(&options).Where("question_id = ?", questionId).Error; err != nil {
		return err
	}

	dbOptions := map[uint]bool{}

	for _, option := range options {
		dbOptions[option.ID] = true
	}

	userOptions := map[uint]bool{}

	for _, option := range optionsUpsert {
		userOptions[option.ID] = true
	}

	// Which options are in DB but not in user options - delete them
	for id := range dbOptions {
		if _, existsInUserOpts := userOptions[id]; !existsInUserOpts {
			if err := dal.db.Delete(&models.Option{}, id).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
