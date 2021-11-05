# Plant Healthcheck Server

### Run

`air .`

### Build and create docker image

`go build -o server`

`docker build . -t sylank/plant-healthcheck-server:latest -f ./Dockerfile`

`docker-compose up`
