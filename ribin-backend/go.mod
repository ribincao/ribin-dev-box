module github.com/ribincao/ribin-dev-box/ribin-backend

go 1.19

require (
	github.com/ribincao/ribin-dev-box/ribin-common v0.0.0-20230629055451-360bbddcf081
	github.com/ribincao/ribin-dev-box/ribin-protocol v0.0.0-20230629053133-4b1e545a0164
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.56.1
)

replace github.com/ribincao/ribin-dev-box/ribin-common => ../ribin-common

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230628200519-e449d1ea0e82 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
