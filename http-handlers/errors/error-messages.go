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
)

func (ErrorMessages) BadRequestMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: msg})
}

func (ErrorMessages) EntityNotFound(ctx *gin.Context, entity string) {
	ErrorMessages.BadRequestMsg(ctx, fmt.Sprintf("%s does not exist.", entity))
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
