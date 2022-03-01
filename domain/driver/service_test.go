package driver

import (
	mocks "bit-driver-location-service/domain/.mocks"
	"bit-driver-location-service/request"
	"bit-driver-location-service/response"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestService_FindNearestDriverLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var repoResponse = []response.DriverLocation{
		{
			Geometry: response.Geometry{
				Type:        "Point",
				Coordinates: []float64{10, 2.21},
			},
		},
		{
			Geometry: response.Geometry{
				Type:        "Point",
				Coordinates: []float64{100, 86.2},
			},
		},
	}
	var loc = request.CustomerLocation{
		Longitude: 100,
		Latitude:  75,
	}

	mockRepo := mocks.NewMockRepositoryImplementation(ctrl)
	mockRepo.
		EXPECT().
		FindNearestDriverByLocation(loc).
		Return(repoResponse, nil).
		Times(1)

	service := NewService(mockRepo, log.Default())

	var actualResponse = service.FindNearestDriverLocation(loc)
	var expectedResponse = []float64{100, 86.2}

	assert.Equal(t, expectedResponse, actualResponse.Geometry.Coordinates)
}
