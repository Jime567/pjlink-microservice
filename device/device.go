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

	route.GET("/raw", d.RawInfo)
	route.POST("/raw", d.Raw)
	route.GET("/command", d.CommandInfo)
	route.POST("/command", d.Command)

	//status endpoints
	router.GET("/:address/power/status", d.GetPowerStatus)
	router.GET("/:address/display/status", d.GetBlankedStatus)
	router.GET("/:address/volume/mute/status", d.GetMuteStatus)
	router.GET("/:address/input/current", d.GetCurrentInput)
	router.GET("/:address/input/list", d.GetInputList)

	//functionality endpoints
	router.GET("/:address/power/on", d.PowerOn)
	router.GET("/:address/power/standby", d.PowerOff)
	router.GET("/:address/display/blank", d.DisplayBlank)
	router.GET("/:address/display/unblank", d.DisplayUnBlank)
	router.GET("/:address/volume/mute", d.VolumeMute)
	router.GET("/:address/volume/unmute", d.VolumeUnMute)
	router.GET("/:address/input/:port", d.SetInputPort)

	server := &http.Server{
		Addr:           port,
		MaxHeaderBytes: 1021 * 10,
	}

	d.Log.Info("running http server", slog.String("port", port))
	err := router.Run(server.Addr)

	d.Log.Error("http server stopped", err)

	return fmt.Errorf("http server stopped")
}
