package router

import (
    "gin-example/controller"
    "github.com/gin-gonic/gin"
)


func SetRouter(router *gin.Engine) {
    api := router.Group("/api")

    api.GET("/hello", controller.Hello)
    api.GET("/albums", controller.GetAlbums)
    api.GET("/albums/:id", controller.GetAlbum)
    api.POST("/albums", controller.PostAlbum)
    api.PUT("/albums/:id", controller.PutAlbum)
    api.DELETE("/albums/:id", controller.DeleteAlbum)
}
