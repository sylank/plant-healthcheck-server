# Plant Healthcheck Server

### Run

`air .`

### Build and create docker image

`go build -o server`

`docker build . -t sylank/plant-healthcheck-server:latest -f ./Dockerfile`

`docker-compose up`

### Example commands

```
curl -d '{"sensor_id":"value1", "command":0, "temperature":1.1, "humidity":2.22, "soil_moisture":3.33}' -H "Content-Type: application/json" -X POST http://localhost:3000/insert
```
