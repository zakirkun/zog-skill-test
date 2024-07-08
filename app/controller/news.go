package controller

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zakirkun/zot-skill-test/app/domain/models"
	"github.com/zakirkun/zot-skill-test/app/domain/types"
	"github.com/zakirkun/zot-skill-test/app/usecase"
	"github.com/zakirkun/zot-skill-test/utils"
)

func CreateNews(ctx *gin.Context) {

	var request types.CreateNews
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	file, _ := ctx.FormFile("thumbnail")
	if file != nil {
		uploadPath := filepath.Join("assets", file.Filename)
		if err := ctx.SaveUploadedFile(file, uploadPath); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
			return
		}

		request.URLThumbnail = uploadPath
	}

	if _, err := usecase.CreateNews(request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.SetGeneralResponse("OK", "Success created", nil))
}

func UpdateNews(ctx *gin.Context) {
	var request types.UpdateNews
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	file, _ := ctx.FormFile("thumbnail")
	if file != nil {
		uploadPath := filepath.Join("assets", file.Filename)
		if err := ctx.SaveUploadedFile(file, uploadPath); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
			return
		}

		request.URLThumbnail = uploadPath
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	if _, err := usecase.UpdateNews(id, request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("OK", "Updated News", nil))
}

func GetNews(ctx *gin.Context) {

	var (
		data *[]models.News
		err  error
	)

	filter := ctx.Query("filter")
	if filter != "" {
		data, err = usecase.FilterNews(filter)
	} else {
		data, err = usecase.ListNews()
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("OK", "Get News", data))
}

func GetDetailNews(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	news, err := usecase.GetDetailNews(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("OK", "Get Detail News", news))
}

func DeleteNews(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	_, err := usecase.DeleteNews(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse("BAD", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SetGeneralResponse("OK", "Delete News", nil))
}
