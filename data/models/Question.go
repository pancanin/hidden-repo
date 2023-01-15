package data

import (
	uuid "github.com/satori/go.uuid"

	"gorm.io/gorm"
)

type QuestionIn struct {
	Body    string     `json:"body" binding:"required,min=10,max=500"`
	Options []OptionIn `json:"options"`
}

type Question struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;"`
	Body    string
	Options []Option
}

type QuestionOut struct {
	ID      uuid.UUID   `json:"id"`
	Body    string      `json:"body"`
	Options []OptionOut `json:"options"`
}

func (q *Question) BeforeCreate(db *gorm.DB) error {
	q.ID = uuid.NewV4()

	return nil
}

func (m *Question) ToResponse() QuestionOut {
	return QuestionOut{
		ID:      m.ID,
		Body:    m.Body,
		Options: ToResponse(m.Options),
	}
}

func ToQuestionsResponse(questions []Question) []QuestionOut {
	var questionsOut []QuestionOut = []QuestionOut{}

	for _, question := range questions {
		questionsOut = append(questionsOut, question.ToResponse())
	}

	return questionsOut
}

func (q *QuestionIn) ToDal() Question {
	return Question{
		Body:    q.Body,
		Options: ToDal(q.Options),
	}
}
