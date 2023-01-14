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

func (dal TodoDal) Create(todoItem *models.TodoIn) (models.TodoModel, error) {
	todo := models.TodoModel{
		Title:     todoItem.Title,
		Completed: todoItem.Completed,
		DueDate:   todoItem.DueDate,
	}

	if err := dal.db.Create(&todo).Error; err != nil {
		return models.TodoModel{}, err
	}

	return todo, nil
}

func (dal TodoDal) GetAll() ([]models.TodoModel, error) {
	var todos []models.TodoModel

	if err := dal.db.Model(models.TodoModel{}).Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func (dal TodoDal) Get(id int) (models.TodoModel, error) {
	var todo models.TodoModel

	if err := dal.db.Model(models.TodoModel{}).First(&todo, id).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (dal TodoDal) Update(todo *models.TodoUpdate) (models.TodoModel, error) {
	err := dal.db.
		Model(models.TodoModel{}).
		Where("id = ?", todo.Id).
		Updates(todo).
		Error

	if err != nil {
		return models.TodoModel{}, err
	}

	return dal.Get(todo.Id)
}

func (dal TodoDal) Delete(id int) error {
	return dal.db.
		Delete(&models.TodoModel{}, id).
		Error
}
