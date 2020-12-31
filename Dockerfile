# build stage
FROM golang:1.15.5-alpine3.12 AS build-env
RUN apk add build-base
ADD . /src
RUN cd /src && go build -o iot-api -i cmd/main.go

# final stage
FROM alpine:3.12
RUN mkdir -p /app/cfg
WORKDIR /app
COPY --from=build-env /src/iot-api /app
ENTRYPOINT [ "./iot-api"]