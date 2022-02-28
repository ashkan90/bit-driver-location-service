package response

type DriverLocation struct {
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}