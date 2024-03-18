package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/byuoitav/pjlink-control/device"
	"github.com/gin-gonic/gin"

	"github.com/spf13/pflag"
)

var logger *slog.Logger

func main() {
	var port, logLev string
	pflag.StringVarP(&port, "port", "p", "8005", "port for microservice to av-api communication")
	pflag.StringVarP(&logLev, "log", "l", "Info", "Initial log level")
	pflag.Parse()

	port = ":" + port
	logLevel := new(slog.LevelVar)

	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	slog.SetDefault(logger)

	manager := device.DeviceManager{
		Log: logger,
	}

	setLogLevel(logLev, logLevel)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "good",
		})
	})

	router.GET("/logLevel/:level", func(context *gin.Context) {
		err := setLogLevel(context.Param("level"), logLevel)
		if err != nil {
			logger.Error("can not set log level", "error", err)
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"current logLevel": logLevel.Level(),
		})
	})

	router.GET("/logLevel", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"current logLevel": logLevel.Level(),
		})
	})

	manager.RunHTTPServer(router, port)
}
