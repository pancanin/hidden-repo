package httphandlers

import (
	"net/http"
	data "questions/data"
	models "questions/data/models"
	httperrors "questions/http-handlers/errors"
	"strconv"

	"github.com/gin-gonic/gin"
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

	// TODO: Add Business logic validation here

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
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handler.httpErrors.BadRequestMsg(ctx, "Invalid id parameter")
		return
	}

	var questionUpdateData models.QuestionIn

	if err := ctx.BindJSON(&questionUpdateData); err != nil {
		handler.httpErrors.BadRequestErr(ctx, err)
		return
	}

	questionToUpdate, err := handler.dal.GetOne(uint(id))

	if err != nil {
		handler.httpErrors.GenericServerErrorEx(ctx, err)
		return
	}

	if questionToUpdate == nil {
		handler.httpErrors.BadRequestMsg(ctx, "Question does not exist.")
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

// func (handler TodoHandler) Delete(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))

// 	if err != nil {
// 		handler.httpErrors.BadRequestMsg(ctx, "Invalid id parameter")
// 		return
// 	}

// 	if _, err := handler.dal.Get(id); err != nil {
// 		handler.httpErrors.BadRequestMsg(ctx, "Todo does not exist")
// 	}

// 	if err := handler.dal.Delete(id); err != nil {
// 		handler.httpErrors.GenericServerErrorEx(ctx, err) // This might not be correct for all cases.
// 		return
// 	}

// 	ctx.Status(http.StatusNoContent)
// }
