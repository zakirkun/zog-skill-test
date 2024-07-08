package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zakirkun/zot-skill-test/app/controller"
)

func NewRouter() http.Handler {

	// Init Router
	r := gin.Default()

	// Cors
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"*"},
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))

	r.MaxMultipartMemory = 8 << 20
	r.Static("/assets", "./assets")

	// Health Checker
	r.GET("/", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong " + fmt.Sprint(time.Now().Unix()),
		})
	})

	api := r.Group("/api")

	topic := api.Group("/topic")
	{
		topic.GET("/", controller.ListTopic)
		topic.GET("/search", controller.SearchTopic)
		topic.GET("/:id", controller.GetDetailTopic)
		topic.POST("/", controller.CreateTopic)
		topic.PATCH("/:id", controller.UpdateTopic)
		topic.DELETE("/:id", controller.DeleteTopic)
	}

	news := api.Group("/news")
	{
		news.GET("/", controller.GetNews)
		news.GET("/:id", controller.GetDetailNews)
		news.POST("/", controller.CreateNews)
		news.PATCH("/:id", controller.UpdateNews)
		news.DELETE("/:id", controller.DeleteNews)
	}

	return r
}
