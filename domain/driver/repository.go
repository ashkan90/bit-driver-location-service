package driver

import (
	"bit-driver-location-service/request"
	"bit-driver-location-service/response"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryDriver struct {
	collection *mongo.Collection
}

func NewRepositoryDriver(collection *mongo.Collection) *RepositoryDriver {
	return &RepositoryDriver{
		collection: collection,
	}
}

func (r *RepositoryDriver) FindNearestDriverByLocation(loc request.CustomerLocation) ([]response.DriverLocation, error) {
	var c, err = r.collection.Find(context.TODO(), bson.M{
		"geometry": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{loc.Longitude, loc.Latitude},
				},
				"$maxDistance": 5000,
			},
		},
	})
	if err != nil {
		return []response.DriverLocation{}, err
	}

	var locations []response.DriverLocation
	err = c.All(context.TODO(), &locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}
