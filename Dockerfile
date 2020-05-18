FROM golang:1.14.3-alpine AS build

WORKDIR /go/src/PrometheusAlert
ADD . .
RUN sh build.sh


FROM alpine:3.9
WORKDIR /app
COPY --from=build /go/src/PrometheusAlert/PrometheusAlert /app
RUN chmod -R 755 /app
CMD ["/app/PrometheusAlert"]
