<h1 align="center">Bi Taksi Location Service</h1>

![example workflow](https://github.com/ashkan90/bit-driver-location-service/actions/workflows/github-actions.yml/badge.svg) _[badge_soon]_


### ✨ [Project Runs On](https://bit-driver-location-service.herokuapp.com/) _soon_

# Hexagonal Architecture
The idea of Hexagonal Architecture is to put inputs and outputs at the edges of our design. Business logic should not depend on whether we expose a REST or a GraphQL API, and it should not depend on where we get data from — a database, a microservice API exposed via gRPC or REST, or just a simple CSV file. <br>

The pattern allows us to isolate the core logic of our application from outside concerns. Having our core logic isolated means we can easily change data source details without a significant impact or major code rewrites to the codebase.

## Run on Local Machine

```shell

```

## Endpoints

```console

```

## Example Usages

```shell
```

## Run Tests

```console
go test ./...
```

## Deployment

> To deploy the case, I used `Heroku`. The deployment is automated with `github actions` and it's `containerized`

## Author

👤 **Emirhan Ataman**


## 📝 License

Copyright © 2021 [Emirhan Ataman](https://github.com/ashkan90). <br />
