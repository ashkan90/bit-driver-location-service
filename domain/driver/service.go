package driver

import (
	"bit-driver-location-service/request"
	"bit-driver-location-service/response"
	"github.com/umahmood/haversine"
)

type Service struct {
	Repository *RepositoryDriver
}

type RepositoryImplementation interface {
	FindNearestDriverByLocation(loc request.CustomerLocation) ([]response.DriverLocation, error)
}

type Distance struct {
	Kilometers float64
	Location   response.DriverLocation
}

func NewService(repo *RepositoryDriver) *Service {
	return &Service{Repository: repo}
}

func (s *Service) FindNearestDriverLocation(loc request.CustomerLocation) response.DriverLocation {
	var drivers, err = s.Repository.FindNearestDriverByLocation(loc)
	if err != nil {
		return response.DriverLocation{}
	}

	var customerCoordinates = haversine.Coord{Lat: loc.Latitude, Lon: loc.Longitude}
	var distances []Distance

	for _, driver := range drivers {
		var driverCoordinates = haversine.Coord{Lat: driver.Geometry.Coordinates[1], Lon: driver.Geometry.Coordinates[0]}
		var _, km = haversine.Distance(customerCoordinates, driverCoordinates)

		distances = append(distances, Distance{
			Kilometers: km,
			Location:   driver,
		})
	}

	var nearest = findMin(distances)

	return nearest.Location
}

func findMin(distances []Distance) Distance {
	var min = distances[0]

	for _, distance := range distances {
		if distance.Kilometers < min.Kilometers {
			min = distance
		}
	}

	return min
}
