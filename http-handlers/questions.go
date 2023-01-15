package httphandlers

import (
	"fmt"
	"net/http"
	data "questions/data"
	models "questions/data/models"
	httperrors "questions/http-handlers/errors"
	validators "questions/validators"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	ID_PARAM_NAME = "id"
)

type QuestionHandler struct {
	dal        *data.QuestionsDal
	httpErrors httperrors.ErrorMessages
}

func NewQuestionHandler(dal *data.QuestionsDal) QuestionHandler {
	return QuestionHandler{dal: dal}
}

func (handler QuestionHandler) Create(ctx *gin.Context) {
	userId, err := uuid.FromString(ctx.Request.Header.Get(models.USER_HEADER_ID))

	if err != nil {
		handler.httpErrors.GenericServerError(ctx)
		return
	}

	var question models.QuestionIn

	if err := ctx.BindJSON(&question); err != nil {
		handler.httpErrors.BadRequestErr(ctx, err)
		return
	}

	if validationMsg := validators.Validate(&question); len(validationMsg) != 0 {
		handler.httpErrors.BadRequestMsg(ctx, validationMsg)
		return
	}

	params := models.QuestionCreateParams{
		Question: question,
		UserID:   userId,
	}

	createdQuestion, err := handler.dal.Create(params)

	if err != nil {
		handler.httpErrors.GenericServerErrorEx(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, createdQuestion.ToResponse())
}

func (handler QuestionHandler) GetPaginated(ctx *gin.Context) {
	userId, err := uuid.FromString(ctx.Request.Header.Get(models.USER_HEADER_ID))

	fmt.Println("Resolved user id: " + userId.String())

	if err != nil {
		handler.httpErrors.GenericServerErrorEx(ctx, err)
		return
	}

	params := models.QuestionsGetPaginatedParams{
		Req:    ctx.Request,
		UserID: userId,
	}

	questions, err := handler.dal.GetPaginated(params)

	if err != nil {
		handler.httpErrors.GenericServerError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, models.ToQuestionsResponse(questions))
}

func (handler QuestionHandler) Update(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param(ID_PARAM_NAME))

	if err != nil {
		handler.httpErrors.BadRequestMsg(ctx, httperrors.INVALID_ID_MSG)
		return
	}

	userId, err := uuid.FromString(ctx.Request.Header.Get(models.USER_HEADER_ID))

	if err != nil {
		handler.httpErrors.GenericServerError(ctx)
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

	getOneParams := models.QuestionGetOneParams{
		QuestionID: id,
		UserID:     userId,
	}

	if _, err := handler.dal.GetOne(getOneParams); err != nil {
		handler.httpErrors.EntityNotFound(ctx, models.QUESTION_ENTITY_NAME)
		return
	}

	updateParams := models.QuestionUpdateParams{
		QuestionID: id,
		Question:   questionUpdateData,
		UserID:     userId,
	}

	if err := handler.dal.Update(updateParams); err != nil {
		handler.httpErrors.GenericServerError(ctx)
		return
	}

	updatedQuestion, err := handler.dal.GetOne(getOneParams)

	if err != nil {
		handler.httpErrors.GenericServerError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, updatedQuestion.ToResponse())
}

func (handler QuestionHandler) Delete(ctx *gin.Context) {
	id, err := uuid.FromString(ctx.Param(ID_PARAM_NAME))

	if err != nil {
		handler.httpErrors.BadRequestMsg(ctx, httperrors.INVALID_ID_MSG)
		return
	}

	userId, err := uuid.FromString(ctx.Request.Header.Get(models.USER_HEADER_ID))

	if err != nil {
		handler.httpErrors.GenericServerError(ctx)
		return
	}

	getOneParams := models.QuestionGetOneParams{
		QuestionID: id,
		UserID:     userId,
	}

	if _, err := handler.dal.GetOne(getOneParams); err != nil {
		handler.httpErrors.EntityNotFound(ctx, models.QUESTION_ENTITY_NAME)
		return
	}

	delParams := models.QuestionDeleteParams{
		QuestionID: id,
		UserID:     userId,
	}

	if err := handler.dal.Delete(delParams); err != nil {
		handler.httpErrors.GenericServerErrorEx(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
