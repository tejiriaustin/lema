package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DefaultStructuredLogs() gin.HandlerFunc {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to init zap logger")
	}
	return structuredLogs(logger)
}

func structuredLogs(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		if c.Writer.Status() >= 500 {
			logger.Error("error",
				zap.String("method", param.Method),
				zap.Int("status_code", param.StatusCode),
				zap.Int("body_size", param.BodySize),
				zap.String("path", param.Path),
				zap.String("method", param.Latency.String()),
			)
		} else {
			logger.Info("message",
				zap.String("method", param.Method),
				zap.Int("status_code", param.StatusCode),
				zap.Int("body_size", param.BodySize),
				zap.String("path", param.Path),
				zap.String("method", param.Latency.String()),
			)
		}
	}
}
