package data

type QuestionIn struct {
	Body    string     `json:"body" binding:"required,min=10,max=500"` // Use template strings here
	Options []OptionIn `json:"options"`
}

type Question struct {
	ID      uint `gorm:"primarykey"`
	Body    string
	Options []Option
}

type QuestionOut struct {
	ID      uint        `json:"id"`
	Body    string      `json:"body"`
	Options []OptionOut `json:"options"`
}

func (m *Question) ToResponse() QuestionOut {
	return QuestionOut{
		ID:      m.ID,
		Body:    m.Body,
		Options: ToResponse(m.Options),
	}
}

func ToQuestionsResponse(questions []Question) []QuestionOut {
	var questionsOut []QuestionOut

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
