package httphandlers

import (
	"net/http"
	"strconv"
	data "questions/data"
	models "questions/data/models"
	httperrors "questions/http-handlers/errors"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	dal        *data.TodoDal
	httpErrors httperrors.ErrorMessages
}

func NewTodoHandler(dal *data.TodoDal) TodoHandler {
	return TodoHandler{dal: dal}
}

func (handler TodoHandler) Create(ctx *gin.Context) {
	var todoItem models.TodoIn

	if err := ctx.BindJSON(&todoItem); err != nil {
		handler.httpErrors.BadRequestErr(ctx, err)
		return
	}

	createdTodoItem, err := handler.dal.Create(&todoItem)

	if err != nil {
		handler.httpErrors.GenericServerErrorEx(ctx, err) // This might not be correct for all cases.
		return
	}

	ctx.JSON(http.StatusCreated, &createdTodoItem)
}

func (handler TodoHandler) GetAll(ctx *gin.Context) {
	todos, err := handler.dal.GetAll()

	if err != nil {
		handler.httpErrors.GenericServerError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, models.ToResponse(todos))
}

func (handler TodoHandler) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handler.httpErrors.BadRequestMsg(ctx, "Invalid id parameter")
		return
	}

	todo, err := handler.dal.Get(id)

	if err != nil {
		handler.httpErrors.GenericServerError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, todo.ToResponse())
}

func (handler TodoHandler) Update(ctx *gin.Context) {
	var todoItem models.TodoUpdate

	if err := ctx.BindJSON(&todoItem); err != nil {
		handler.httpErrors.BadRequestErr(ctx, err)
		return
	}

	updatedTodoItem, err := handler.dal.Update(&todoItem)

	if err != nil {
		handler.httpErrors.GenericServerError(ctx) // This might not be correct for all cases.
		return
	}

	ctx.JSON(http.StatusOK, updatedTodoItem.ToResponse())
}

func (handler TodoHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		handler.httpErrors.BadRequestMsg(ctx, "Invalid id parameter")
		return
	}

	if _, err := handler.dal.Get(id); err != nil {
		handler.httpErrors.BadRequestMsg(ctx, "Todo does not exist")
	}

	if err := handler.dal.Delete(id); err != nil {
		handler.httpErrors.GenericServerErrorEx(ctx, err) // This might not be correct for all cases.
		return
	}

	ctx.Status(http.StatusNoContent)
}
