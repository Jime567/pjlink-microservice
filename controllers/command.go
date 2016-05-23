package controllers

import (
	//"fmt"
	"net/http"

	"github.com/byuoitav/pjlink-service/helpers"

	"github.com/labstack/echo"
)

func Test(context echo.Context) error {
	response, err := helpers.Test()
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, response)
}

func PjlinkRequest(context echo.Context) error {
	parsedResponse, err := helpers.PjlinkRequest(context.Param("address"),
		context.Param("port"), context.Param("class"), context.Param("passwd"),
		context.Param("command"), context.Param("param"))

	if err != nil {
		//TODO
		return err
	}

	return context.JSON(http.StatusOK, parsedResponse)
}