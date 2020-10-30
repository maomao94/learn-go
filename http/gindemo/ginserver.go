package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	const keyRequestId = "requestId"
	r.Use(func(context *gin.Context) {
		s := time.Now()

		context.Next()

		// path,status,log latency
		logger.Info("incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))
	}, func(context *gin.Context) {
		context.Set(keyRequestId, rand.Int())
		context.Next()
	})
	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message":    "pong",
			keyRequestId: 1234,
		}
		if rid, exists := c.Get(keyRequestId); exists {
			h[keyRequestId] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
