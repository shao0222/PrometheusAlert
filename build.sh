set GO11MODULE=on
set GO111MODULE=on
go mod init PrometheusAlert
go mod vendor
CGO_ENABLED=0 go build