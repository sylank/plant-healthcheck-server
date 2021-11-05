FROM golang:1.16
WORKDIR /app/plant-healthcheck-server
ADD . /app/plant-healthcheck-server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-X main.VERSION=docker' -o server

FROM debian:stretch
WORKDIR /app
RUN mkdir -p /app/public
RUN apt-get update && apt-get install -y ca-certificates
VOLUME /app/public
ADD public ./public
COPY --from=0 /app/plant-healthcheck-server/server .
EXPOSE 3333
EXPOSE 1323
ENTRYPOINT [ "/app/server" ]
CMD []