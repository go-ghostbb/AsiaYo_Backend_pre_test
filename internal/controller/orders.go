package controller

import (
	"AsiaYo-Backend-pre-test/internal/model"
	"AsiaYo-Backend-pre-test/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"

	. "AsiaYo-Backend-pre-test/internal/utils/response"
)

func CreateOrder(c *gin.Context) {
	var (
		in  model.CreateOrder
		err error
	)

	// 解析並檢查
	if err = c.ShouldBindJSON(&in); err != nil {
		Responder(Mount(c)).FailWithMsg(http.StatusBadRequest, err.Error(), nil)
		return
	}

	// service
	err = service.CreateOrder(&in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(http.StatusBadRequest, err.Error(), in)
		return
	}

	// ok
	Responder(Mount(c)).Ok(in)
}
