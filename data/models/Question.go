package data

import (
	"gorm.io/gorm"
)

type QuestionIn struct {
	Body    string     `json:"body" binding:"required"`
	Options []OptionIn `json:"options"`
}

type OptionIn struct {
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
}

type Question struct {
	gorm.Model
	Body    string   `json:"body" binding:"required"`
	Options []Option `json:"options"`
}

type Option struct {
	gorm.Model
	Body       string `json:"body"`
	Correct    bool   `json:"correct"`
	QuestionID uint
}

type QuestionOut struct {
	ID      uint        `json:"id"`
	Body    string      `json:"body" binding:"required"`
	Options []OptionOut `json:"options"`
}

type OptionOut struct {
	ID      uint   `json:"id"`
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
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
