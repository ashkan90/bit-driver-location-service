package handler

import (
	"bit-driver-location-service/adapters/validator"
	mocks "bit-driver-location-service/domain/.mocks"
	"bit-driver-location-service/domain/driver"
	"bit-driver-location-service/request"
	"bit-driver-location-service/response"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAsd(t *testing.T) {
	var loc = request.CustomerLocation{
		Longitude: 100,
		Latitude:  75,
	}
	var body = fmt.Sprintf(`{"longitude": %f, "latitude": %f}`, loc.Longitude, loc.Latitude)

	req := httptest.NewRequest(http.MethodGet, "/nearest-driver-location", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	e.Validator = validator.NewRequestValidator()
	ctx := e.NewContext(req, rec)

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
	var expectedResponse = `{"coordinates":[100,86.2]}` + "\n"

	mockRepo := mocks.NewMockRepositoryImplementation(ctrl)
	mockRepo.
		EXPECT().
		FindNearestDriverByLocation(loc).
		Return(repoResponse, nil).
		Times(1)

	handler := &DriverHandler{
		Service: driver.NewService(mockRepo, log.Default()),
	}

	err := handler.FindDriverLocation(ctx)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, rec.Body.String())
}
