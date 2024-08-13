package router

import (
	"AsiaYo-Backend-pre-test/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitOrder(router *gin.RouterGroup) {
	router.POST("", controller.CreateOrder)
}
