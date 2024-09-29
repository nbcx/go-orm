module github.com/nbcx/go-orm

go 1.23.0

replace github.com/nbcx/log => ../log

require (
	github.com/beego/beego/v2 v2.3.1
	github.com/go-sql-driver/mysql v1.8.1
	github.com/hashicorp/golang-lru v1.0.2
	github.com/lib/pq v1.10.9
	github.com/mattn/go-sqlite3 v1.14.23
	github.com/nbcx/log v0.0.0-20240929031720-b28ab0992f48
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.20.4
	github.com/stretchr/testify v1.9.0
	github.com/valyala/bytebufferpool v1.0.0
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.30.0
	go.opentelemetry.io/otel/sdk v1.30.0
	go.opentelemetry.io/otel/trace v1.30.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20230509054315-a9deabde6e02 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
