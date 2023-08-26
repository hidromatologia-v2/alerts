FROM golang:1.21-alpine AS build-stage

RUN apk add upx
WORKDIR /alerts-src
COPY . .
RUN go build -o /alerts .
RUN upx /alerts

FROM alpine:latest AS release-stage

COPY --from=build-stage /alerts /alerts
# -- Environment variables
ENV MEMPHIS_CONSUMER_STATION     "alerts"
ENV MEMPHIS_CONSUMER_CONSUMER    "alerts-consumer"
ENV MEMPHIS_CONSUMER_HOST        "memphis"
ENV MEMPHIS_CONSUMER_USERNAME    "root"
ENV MEMPHIS_CONSUMER_PASSWORD    "memphis"
ENV MEMPHIS_CONSUMER_CONN_TOKEN  ""
ENV MEMPHIS_PRODUCER_STATION     "messages"
ENV MEMPHIS_PRODUCER_PRODUCER    "messages-producer"
ENV MEMPHIS_PRODUCER_HOST        "memphis"
ENV MEMPHIS_PRODUCER_USERNAME    "root"
ENV MEMPHIS_PRODUCER_PASSWORD    "memphis"
ENV MEMPHIS_PRODUCER_CONN_TOKEN  ""
ENV POSTGRES_DSN                 "host=postgres user=sulcud password=sulcud dbname=sulcud port=5432 sslmode=disable"
# -- Environment variables
ENTRYPOINT [ "sh", "-c", "/alerts" ]