package httphandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorMessages struct{}

type ErrorResponse struct {
	Message string `json:"message"`
}

const (
	INTERNAL_SERVER_ERROR_MSG = "We are experiencing issues at the moment. Please try again later."
	INVALID_ID_MSG            = "Invalid id parameter"
	ENTITY_MISSING_MSG        = "%s does not exist."
	INVALID_PAGINATION_PARAMS = "Invalid pagination params. Both 'page' and 'pageSize' are required. Omit both to disable pagination."
)

func (ErrorMessages) BadRequestMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: msg})
}

func (em ErrorMessages) EntityNotFound(ctx *gin.Context, entity string) {
	em.BadRequestMsg(ctx, fmt.Sprintf(ENTITY_MISSING_MSG, entity))
}

func (ErrorMessages) BadRequestErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
}

func (ErrorMessages) GenericServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: INTERNAL_SERVER_ERROR_MSG})
}

func (ErrorMessages) GenericServerErrorEx(ctx *gin.Context, e error) {
	ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: e.Error()})
}
