package validators

import (
	models "questions/data/models"
)

const (
	OPTIONS_COUNT_MSG        = "Question should have more than one answer."
	CORRECT_ANSWER_COUNT_MSG = "Question should have at least one correct answer."
	UNIQUE_OPTION_MSG        = "Question should have unique answers."
)

func Validate(q *models.QuestionIn) string {
	if len(q.Options) <= 1 {
		return OPTIONS_COUNT_MSG
	}

	/* At least one correct answer */
	correctAnswersCount := 0

	for _, option := range q.Options {
		if option.Correct {
			correctAnswersCount += 1
		}
	}

	if correctAnswersCount == 0 {
		return CORRECT_ANSWER_COUNT_MSG
	}

	/* Unique options */
	uniqueOptions := map[string]bool{}

	for _, option := range q.Options {
		uniqueOptions[option.Body] = true
	}

	if len(uniqueOptions) != len(q.Options) {
		return UNIQUE_OPTION_MSG
	}

	return ""
}
