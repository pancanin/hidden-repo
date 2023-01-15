package validators

import (
	models "questions/data/models"
)

func Validate(q *models.QuestionIn) string {
	if len(q.Options) <= 1 {
		return "Question should have more than one answer."
	}

	uniqueOptions := map[string]bool{}

	for _, option := range q.Options {
		uniqueOptions[option.Body] = true
	}

	if len(uniqueOptions) != len(q.Options) {
		return "Question should have unique answers."
	}

	return ""
}
