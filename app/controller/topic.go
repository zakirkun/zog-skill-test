package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zakirkun/zot-skill-test/app/domain/types"
	"github.com/zakirkun/zot-skill-test/app/usecase"
	"github.com/zakirkun/zot-skill-test/utils"
)

func CreateTopic(ctx *gin.Context) {
	var request types.CreateTopic
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	if _, err := usecase.CreateTopic(request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SetGeneralResponse("OK", "Success created", nil))
}

func UpdateTopic(ctx *gin.Context) {
	var request types.UpdateeTopic
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	if _, err := usecase.UpdateTopic(id, request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("OK", "Success updated", nil))
}

func ListTopic(ctx *gin.Context) {
	topics, err := usecase.ListTopic()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("OK", "Topic loaded", topics))
}

func SearchTopic(ctx *gin.Context) {

	keyword := ctx.Query("q")
	topics, err := usecase.SearchTopic(keyword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("OK", "Topic loaded", topics))
}

func GetDetailTopic(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	topic, err := usecase.GetDetailTopic(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("OK", "Topic loaded", topic))
}

func DeleteTopic(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	topic, err := usecase.DeleteTopic(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("OK", "Topic Deleted", topic))
}
