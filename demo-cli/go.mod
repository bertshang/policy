module github.com/bertshang/policy/demo-cli

go 1.13

require (
	github.com/bertshang/policy/common v0.0.0-20210112062123-6f5bd02d9bef
	github.com/bertshang/policy/demo-service v0.0.0-20210111084732-c3ee1b924881
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/wrapper/breaker/hystrix v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd
	github.com/opentracing/opentracing-go v1.2.0
)
