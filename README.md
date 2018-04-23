# ingresse-backend

ingresse-backend is a simple application developed as part of the selection process for [Ingresse](https://www.ingresse.com/).
The application was built using [Go](https://golang.org/), and the [Macaron](https://go-macaron.com/) framework.

ingresse-backend is a REST API with only one endpoint: `/users`. It implements the http verbs `POST`, `GET`, `PUT` and `DELETE`.

ingresse-backend uses [MongoDB](https://www.mongodb.com) for data persistence.

## Setting up

Since Go is not a major programming language yet, not everyone has a set up environment for it, so two options are available:
- Run natively (requires Go environment installed)
- Run under [Docker](https://www.docker.com/) (visual representation of code coverage not available)

### Running natively
Assuming a working Go environment, run:
```
go get -v github.com/PhillBeck/ingresse-backend/...
cd $GOPATH/src/github/PhillBeck/ingresse-backend
```
Then you have to set up the infrastructure. As mentioned above, ingresse-backend requires a MongoDB instance running on port 27017. For ease of use, a docker-compose file is provided to set up Mongo:
```
cd docker/infrastructure
docker-compose up
```
after that, just navigate to the directory and execute the program:
```
cd $GOPATH/src/github/PhillBeck/ingresse-backend
go run main.go
```

### Running under Docker
This assumes you have both docker AND docker-compose installed.

To run under Docker, simply clone the repository:
```
git clone https://github.com/PhillBeck/ingresse-backend.git
```

then run: 
```
cd ingresse-backend/docker/application
docker-compose up
```
Using this method, the infrastructure is set up automatically.

## Documentation

The repository contains the swagger file in both json and yaml formats.
To access the documentation in a more friendly way, a html version is also available on the `docs` folder.

If you are running the application under Docker, just navigate to `localhost:8000` to view the docs.

If you are running natively, you can run:
```
cd $GOPATH/src/github/PhillBeck/ingresse-backend/docker/docs
docker-compose up
```
and then access `localhost:8000`.

## Testing

### Natively
Assuming a working Go environment, just run:
```
cd $GOPATH/src/github/PhillBeck/ingresse-backend
go test -coverprofile=cover.prof ./...
```
`go test` will output the summary of the tests to stdout. This includes code coverage for each package. For a visual representation of the code coverage, run:
```
go tool cover -html=cover.prof
```

### Under Docker
After the container is running, run:
```
docker exec -it ingresse-backend /bin/sh
go test -coverprofile=cover.prof ./...
```
Unfortunately, the visual representation of the code coverage is not available when running under Docker.