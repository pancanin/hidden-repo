package httphandlers

import (
	"net/http"
	data "questions/data"
	models "questions/data/models"
	httperrors "questions/http-handlers/errors"
	validators "questions/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	ID_PARAM_NAME        = "id"
	QUESTION_ENTITY_NAME = "Question"
)

type QuestionHandler struct {
	dal        *data.QuestionsDal
	httpErrors httperrors.ErrorMessages
}

func NewQuestionHandler(dal *data.QuestionsDal) QuestionHandler {
	return QuestionHandler{dal: dal}
}

func (handler QuestionHandler) Create(ctx *gin.Context) {
	var question models.QuestionIn

	if err := ctx.BindJSON(&question); err != nil {
		handler.httpErrors.BadRequestErr(ctx, err)
		return
	}

	if validationMsg := validators.Validate(&question); len(validationMsg) != 0 {
		handler.httpErrors.BadRequestMsg(ctx, validationMsg)
		return
	}

	createdQuestion, err := handler.dal.Create(&question)

	if err != nil {
		handler.httpErrors.GenericServerErrorEx(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdQuestion.ToResponse())
}

func (handler QuestionHandler) GetAll(ctx *gin.Context) {
	questions, err := handler.dal.GetAll()

	if err != nil {
		handler.httpErrors.GenericServerError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, models.ToQuestionsResponse(questions))
}

func (handler QuestionHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param(ID_PARAM_NAME))

	if err != nil {
		handler.httpErrors.BadRequestMsg(ctx, httperrors.INVALID_ID_MSG)
		return
	}

	var questionUpdateData models.QuestionIn

	if err := ctx.BindJSON(&questionUpdateData); err != nil {
		handler.httpErrors.BadRequestErr(ctx, err)
		return
	}

	if validationMsg := validators.Validate(&questionUpdateData); len(validationMsg) != 0 {
		handler.httpErrors.BadRequestMsg(ctx, validationMsg)
		return
	}

	if _, err := handler.dal.GetOne(uint(id)); err != nil {
		handler.httpErrors.EntityNotFound(ctx, QUESTION_ENTITY_NAME)
		return
	}

	if err := handler.dal.Update(uint(id), &questionUpdateData); err != nil {
		handler.httpErrors.GenericServerError(ctx)
		return
	}

	updatedQuestion, err := handler.dal.GetOne(uint(id))

	if err != nil {
		handler.httpErrors.GenericServerError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, updatedQuestion.ToResponse())
}
