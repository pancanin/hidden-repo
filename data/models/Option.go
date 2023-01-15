package data

import uuid "github.com/satori/go.uuid"

type OptionIn struct {
	Body    string `json:"body" binding:"required,min=1,max=2000"`
	Correct bool   `json:"correct"`
}

type Option struct {
	Body       string
	Correct    bool
	QuestionID uuid.UUID
}

type OptionOut struct {
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
}

func (o *Option) ToResponse() OptionOut {
	return OptionOut{
		Body:    o.Body,
		Correct: o.Correct,
	}
}

func ToResponse(options []Option) []OptionOut {
	var optionsOut []OptionOut = []OptionOut{}

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
