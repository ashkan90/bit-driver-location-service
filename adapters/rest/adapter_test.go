package rest

import (
	"bit-driver-location-service/config"
	"log"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAdapter_ServeShouldAddRespectiveRoutesCorrectly(t *testing.T) {
	var adapter = &Adapter{
		Config: &config.GeneralConfig{
			Server: config.Server{
				Port: "1111",
			},
			Database: config.Database{
				DSN: "mongodb://dsn",
			},
		},
		Logger: log.Default(),
		Server: echo.New(),
	}

	adapter.Serve(nil)

	assert.Len(t, adapter.Server.Routes(), 1)
}
