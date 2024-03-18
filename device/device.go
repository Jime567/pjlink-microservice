package device

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceManager struct {
	Log *slog.Logger
}

func (d *DeviceManager) GetLogger() *slog.Logger {
	return d.Log
}

func (d *DeviceManager) RunHTTPServer(router *gin.Engine, port string) error {
	d.Log.Info("registering http endpoints")

	route := router.Group("")

	route.GET("/raw", handlers.RawInfo)
	route.POST("/raw", handlers.Raw)
	route.GET("/command", handlers.CommandInfo)
	route.POST("/command", handlers.Command)

	//status endpoints
	router.GET("/:address/power/status", handlers.GetPowerStatus)
	router.GET("/:address/display/status", handlers.GetBlankedStatus)
	router.GET("/:address/volume/mute/status", handlers.GetMuteStatus)
	router.GET("/:address/input/current", handlers.GetCurrentInput)
	router.GET("/:address/input/list", handlers.GetInputList)

	//functionality endpoints
	router.GET("/:address/power/on", handlers.PowerOn)
	router.GET("/:address/power/standby", handlers.PowerOff)
	router.GET("/:address/display/blank", handlers.DisplayBlank)
	router.GET("/:address/display/unblank", handlers.DisplayUnBlank)
	router.GET("/:address/volume/mute", handlers.VolumeMute)
	router.GET("/:address/volume/unmute", handlers.VolumeUnMute)
	router.GET("/:address/input/:port", handlers.SetInputPort)

	server := &http.Server{
		Addr:           port,
		MaxHeaderBytes: 1021 * 10,
	}

	d.Log.Info("running http server", slog.String("port", port))
	err := router.Run(server.Addr)

	d.Log.Error("http server stopped", err)

	return fmt.Errorf("http server stopped")
}
