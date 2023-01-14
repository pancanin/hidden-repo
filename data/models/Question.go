package data

type QuestionIn struct {
	Body    string     `json:"body" binding:"required,min=10,max=500"` // Use template strings here
	Options []OptionIn `json:"options"`
}

type OptionIn struct {
	Body    string `json:"body" binding:"required,min=1,max=2000"`
	Correct bool   `json:"correct"`
}

type Question struct {
	ID      uint     `gorm:"primarykey"`
	Body    string   `json:"body"`
	Options []Option `json:"options"`
}

type Option struct {
	ID         uint   `gorm:"primarykey"`
	Body       string `json:"body"`
	Correct    bool   `json:"correct"`
	QuestionID uint
}

type QuestionOut struct {
	ID      uint        `json:"id"`
	Body    string      `json:"body"`
	Options []OptionOut `json:"options"`
}

type OptionOut struct {
	ID      uint   `json:"id"`
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
}

type OptionUpsert struct { // this should be called upsert
	ID      uint   `json:"id"`
	Body    string `json:"body" binding:"required,min=1,max=500"`
	Correct bool   `json:"correct"`
}

type QuestionUpdate struct {
	Body    string         `json:"body" binding:"required,min=10,max=2000"`
	Options []OptionUpsert `json:"options"`
}

func (m *Question) ToResponse() QuestionOut {
	return QuestionOut{
		ID:      m.ID,
		Body:    m.Body,
		Options: ToResponse(m.Options),
	}
}

func (o *Option) ToResponse() OptionOut {
	return OptionOut{
		ID:      o.ID,
		Body:    o.Body,
		Correct: o.Correct,
	}
}

func ToResponse(options []Option) []OptionOut {
	var optionsOut []OptionOut

	for _, option := range options {
		optionsOut = append(optionsOut, option.ToResponse())
	}

	return optionsOut
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

func (o *OptionIn) ToDal() Option {
	return Option{
		Body:    o.Body,
		Correct: o.Correct,
	}
}

func ToDal(optionsIn []OptionIn) []Option {
	var options []Option

	for _, option := range optionsIn {
		options = append(options, option.ToDal())
	}

	return options
}

func (q *QuestionUpdate) ToDal() Question {
	return Question{
		Body: q.Body,
	}
}

func (o *OptionUpsert) ToDal(questionId uint) Option {
	return Option{
		Body:       o.Body,
		Correct:    o.Correct,
		QuestionID: questionId,
	}
}
