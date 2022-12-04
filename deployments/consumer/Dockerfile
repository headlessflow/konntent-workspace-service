FROM golang:1.18-alpine as build

WORKDIR /consumer

VOLUME ["/consumer"]

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY .. .

RUN GOOS=linux CGO_ENABLED=0 go build -o /konntent-service-consumer-service ./cmd/consumer

FROM alpine
COPY --from=build /konntent-service-consumer-service ./consumer
COPY --from=build ./.env ./.env

COPY ./server.yaml ./
EXPOSE 3000

CMD [ "./consumer" ]