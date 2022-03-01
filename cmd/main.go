package main

import (
	"bit-driver-location-service/adapters/database/mongo"
	"bit-driver-location-service/adapters/rest"
	"bit-driver-location-service/config"
	"context"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Rest adapter which serves the routes via given settings
	var confPath string
	var port = os.Getenv("PORT")
	var logger = log.Default()
	var sv = NewServer()

	flag.StringVar(&confPath, "service-config", "service_config.yaml", "")
	flag.Parse()

	var conf = config.NewGeneralConfig(confPath)
	log.Println(conf)

	if port != "" {
		conf.Server.Port = port
	}

	var locationsCollStr = "locations"
	var locationsColl = mongo.Connect(conf.Database.DSN, locationsCollStr).Collection(locationsCollStr)

	var restAdapter = &rest.Adapter{
		Config: &conf,
		Logger: logger,
		Server: sv,
	}
	restAdapter.Serve(locationsColl)
	go func() {
		logger.Fatal(sv.Start(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)))
	}()

	gracefulShutdown(logger, sv)
}

func gracefulShutdown(logger *log.Logger, s *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}
}
