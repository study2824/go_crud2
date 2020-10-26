package server

import (
	"github.com/gin-gonic/gin"
	"github.com/koko2824/goSample/controllers"
)

func Init()  {
	r := router()
	r.Run("localhost:8080")
}

func router() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	ctrl := controllers.Controller{}
	r.GET("/", ctrl.Index)
	r.GET("/add", ctrl.AddForm)
	r.POST("/create", ctrl.Create)
	r.POST("/update/:id", ctrl.Update)
	r.GET("/detail/:id", ctrl.Detail)
	r.POST("/delete/:id", ctrl.Delete)

	return r
}

