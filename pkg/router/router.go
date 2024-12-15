package router

import (
	"main/pkg/appLogs"
	"main/pkg/endpoints/userHandler"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewServer() *gin.Engine {
	router := gin.Default()
	router.Use(Middleware(appLogs.ApplicationLog))
	addRoutes(router)
	return router
}

func Middleware(appLogger *zap.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()

		// Log request details
		duration := time.Since(startTime)
		appLogger.Info("HTTP request",
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Duration("latency", duration),
		)
	}
}

// func timeLog(c *gin.Context) {
// 	fmt.Println("Request start ....")
// 	timeStart := time.Now()
// 	c.Next()
// 	latency := time.Since(timeStart)
// 	fmt.Printf("API time: %v \n", latency)
// }

func addRoutes(r *gin.Engine) {
	userValidator := userHandler.NewValidator()
	userHandlerInstance := userHandler.NewHandler(userValidator)
	r.POST("/check", userHandlerInstance.AddUser)
}
