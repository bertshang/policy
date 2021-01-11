build:
	protoc -I. --go_out=plugins=micro:. proto/demo/demo.proto
	docker build -t policy-demo-service .

run:
	docker run -d -p 9091:9091 -e MICRO_SERVER_ADDRESS=:9091 -e MICRO_REGISTRY=mdns policy-demo-service