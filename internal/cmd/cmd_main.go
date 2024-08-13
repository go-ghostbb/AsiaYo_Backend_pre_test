package cmd

import (
	"AsiaYo-Backend-pre-test/internal/router"
	"context"
	"github.com/gin-gonic/gin"
)

func MainFn(ctx context.Context) (err error) {
	engine := gin.Default()

	// order
	router.InitOrder(engine.Group("/api/orders"))

	err = engine.Run(":8080")
	return
}
