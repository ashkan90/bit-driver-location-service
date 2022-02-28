package handler

import (
	"bit-driver-location-service/domain/driver"
	"bit-driver-location-service/request"
	"bit-driver-location-service/response"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	ErrorOnBind = `{"message": "%s"}`
)

type DriverHandler struct {
	Service *driver.Service
}

type DriverImplementations interface {
	FindNearestDriverLocation(loc request.CustomerLocation) response.DriverLocation
}

func (h *DriverHandler) FindDriverLocation(c echo.Context) error {
	var err error
	var cLocation request.CustomerLocation
	if err = c.Bind(&cLocation); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf(ErrorOnBind, err.Error()))
	}

	if err = c.Validate(&cLocation); err != nil {
		return err
	}

	var driverLocation = h.Service.FindNearestDriverLocation(cLocation)

	return c.JSON(http.StatusOK, driverLocation.Geometry)
}
