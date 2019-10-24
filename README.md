# API for Vigilante Web Heist

## API Endpoints

These are the endpoints that the API currently supports. If any of these
endpoints are out of date or not working, please create an issue.

Basing off [this article](https://restfulapi.net/resource-naming/) on REST API
naming conventions:

- GET /api/honeyclients (**get all honeyclients, or maybe first 50 or
  something**)
- POST /api/honeyclients (**creates a honeyclient**)
- GET /api/honeyclients/{id} (**get all information associated with a specific
  honeyclient**)
- GET /api/honeyclients/{id}/artifacts (**get all file artifacts like
  screenshots or js associated with specific honeyclient**)
- GET /api/honeyclients/{id}/artifacts/{artifact-name} (**get specific artifact
  associated with a honeyclient**)

## How to run

Prerequisites:

* Golang is installed
* Docker is installed

#### Steps to run:

**With Golang**
1. Run `go build` to compile the executable from the file "main.go"
2. Run the executable named "api" (ex: if using mac run `./api`)

**With Docker**
1. Run `docker run -p 8080:8080 csci4950tgt/api`

#### Steps to rebuild Dockerfile and push new image to Docker repo

1. Run `docker build -t csci4950tgt/api .` to build the image
2. Run `docker push csci4950tgt/api` to push new image to Docker repo

#### Steps to test:

1. Run `go test` (or `go test -v` if you want verbose information)
