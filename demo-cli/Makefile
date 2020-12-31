build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t policy-demo-cli .

run:
	docker run policy-demo-cli