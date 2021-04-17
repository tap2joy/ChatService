module github.com/tap2joy/ChatService

go 1.14

require (
	cloud.google.com/go v0.56.0 // indirect
	github.com/tap2joy/Protocols v0.0.0-00010101000000-000000000000
	github.com/go-xorm/xorm v0.7.9
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4
	github.com/lib/pq v1.10.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/procfs v0.0.8 // indirect
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/spf13/viper v1.7.1
	go.elastic.co/apm/module/apmgrpc v1.11.0
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	google.golang.org/grpc v1.37.0
	google.golang.org/grpc/examples v0.0.0-20210409234925-fab5982df20a // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/tap2joy/Protocols => ./proto