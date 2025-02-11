package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/byuoitav/common/status"
	"github.com/byuoitav/pjlink-microservice/pjlink"
	"github.com/gin-gonic/gin"
	"github.com/jessemillar/jsonresp"
)

func (d *DeviceManager) PowerOn(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "power", "power-on")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, status.Power{"on"})
	}

	return context.JSON(http.StatusInternalServerError, response)
}

func (d *DeviceManager) PowerOff(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "power", "power-off")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, status.Power{"standby"})
	}

	return context.JSON(http.StatusOK, response)
}

func (d *DeviceManager) GetPowerStatus(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "power", "query")

	response, err := pjlink.GetPowerStatus(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

// some projectors *panasonic - cough* only accept av mute, not just blank, so
// a blank command both blanks and mutes
func (d *DeviceManager) DisplayBlank(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "av-mute-on")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, status.Blanked{true})
	}

	return context.JSON(http.StatusOK, response)
}

func (d *DeviceManager) DisplayUnBlank(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "av-mute-off")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, status.Blanked{false})
	}

	return context.JSON(http.StatusOK, response)
}

func (d *DeviceManager) GetBlankedStatus(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "query")

	response, err := pjlink.GetBlankedStatus(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (d *DeviceManager) VolumeMute(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "audio-mute-on")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, status.Mute{true})
	}

	return context.JSON(http.StatusOK, response)
}

func (d *DeviceManager) VolumeUnMute(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "audio-mute-off")

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, status.Mute{false})
	}

	return context.JSON(http.StatusOK, response)
}

func (d *DeviceManager) GetMuteStatus(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "av-mute", "query")

	response, err := pjlink.GetMuteStatus(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (d *DeviceManager) SetInputPort(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "input", context.Param("port"))

	response, responseError := pjlink.HandleRequest(request)
	if responseError != nil {
		jsonresp.New(context.Response(), http.StatusBadRequest, responseError.Error())
		return nil
	}

	if contains("success", response.Response) || contains("OK", response.Response) {
		return context.JSON(http.StatusOK, status.Input{context.Param("port")})
	}

	return context.JSON(http.StatusInternalServerError, response)
}

func (d *DeviceManager) GetCurrentInput(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "input", "query")

	response, err := pjlink.GetCurrentInput(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func (d *DeviceManager) GetInputList(context *gin.Context) error {
	request := formRequestFromEnvVars(context.Param("address"), "input-list", "query")

	response, err := pjlink.GetInputList(request)
	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, response)
}

func formRequestFromEnvVars(address, command, parameter string) pjlink.PJRequest {
	request := pjlink.PJRequest{
		Address:   address,
		Port:      os.Getenv("PJLINK_PORT"),
		Password:  os.Getenv("PJLINK_PASS"),
		Class:     "1",
		Command:   command,
		Parameter: parameter,
	}

	return request
}

func contains(str string, list []string) bool {

	for _, s := range list {

		if strings.Contains(s, str) {
			return true
		}
	}

	return false
}
