FROM golang:1.17-alpine as build

WORKDIR /app

VOLUME ["/app"]

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o /bit-driver-location-service ./cmd/

FROM alpine
COPY --from=build /bit-driver-location-service ./app

COPY ./service_config.yaml coordinates.geojson import.sh ./
EXPOSE 8080

CMD [ "./app", "--service-config", "service_config.yaml" ]