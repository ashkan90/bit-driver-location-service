package rest

import (
	"bit-driver-location-service/adapters/handler"
	"bit-driver-location-service/config"
	"bit-driver-location-service/domain/driver"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Adapter struct {
	Config *config.GeneralConfig
	Logger *log.Logger
	Server *echo.Echo
}

func (a *Adapter) Serve(coll *mongo.Collection) {
	var driverRepository = driver.NewRepositoryDriver(coll)
	var driverService = driver.NewService(driverRepository, a.Logger)
	var driverRest = &handler.DriverHandler{Service: driverService}

	a.Server.Add(http.MethodGet, "/nearest-driver-location", driverRest.FindDriverLocation)

	a.Logger.Println("Server will be started on port " + a.Config.Server.Port)
}
