package httphandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorMessages struct{}

type ErrorResponse struct {
	Message string `json:"message"`
}

const (
	internalServerErrorMessage = "We are experiencing issues at the moment. Please try again later."
)

func (ErrorMessages) BadRequestMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: msg})
}

func (ErrorMessages) BadRequestErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
}

func (ErrorMessages) GenericServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: internalServerErrorMessage})
}

func (ErrorMessages) GenericServerErrorEx(ctx *gin.Context, e error) {
	ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: e.Error()})
}
