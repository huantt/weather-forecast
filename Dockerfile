FROM golang:1.21-bullseye as builder

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y tzdata ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/app .
COPY ./template ./template
RUN ln -s /app/app /bin/weather-forecast
ENTRYPOINT ["weather-forecast", "update-weather"]