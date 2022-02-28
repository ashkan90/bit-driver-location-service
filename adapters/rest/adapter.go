package rest

import (
	"bit-driver-location-service/adapters/database/mongo"
	"bit-driver-location-service/adapters/handler"
	"bit-driver-location-service/config"
	"bit-driver-location-service/domain/driver"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Adapter struct {
	Config config.Server
	Logger *log.Logger
	Server *echo.Echo
}

func (a *Adapter) Serve() {
	var locationsColl = mongo.Connect().Collection("locations")
	var driverRepository = driver.NewRepositoryDriver(locationsColl)
	var driverService = driver.NewService(driverRepository)
	var driverRest = &handler.DriverHandler{Service: driverService}

	a.Server.Add(http.MethodGet, "/nearest-driver-location", driverRest.FindDriverLocation)

	a.Logger.Println("Server has started on port " + a.Config.Port)
	a.Logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", a.Config.Host, a.Config.Port), a.Server))
}
