package rest

import (
	"bit-driver-location-service/adapters/database/mongo"
	"bit-driver-location-service/adapters/handler"
	"bit-driver-location-service/config"
	"bit-driver-location-service/domain/driver"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Adapter struct {
	Config *config.GeneralConfig
	Logger *log.Logger
	Server *echo.Echo
}

func (a *Adapter) Serve() {
	var locationsCollStr = "locations"
	var locationsColl = mongo.Connect(a.Config.Database.DSN, locationsCollStr).Collection(locationsCollStr)
	var driverRepository = driver.NewRepositoryDriver(locationsColl)
	var driverService = driver.NewService(driverRepository, a.Logger)
	var driverRest = &handler.DriverHandler{Service: driverService}

	a.Server.Add(http.MethodGet, "/nearest-driver-location", driverRest.FindDriverLocation)

	a.Logger.Println("Server has started on port " + a.Config.Server.Port)
	a.Logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", a.Config.Server.Host, a.Config.Server.Port), a.Server))
}
