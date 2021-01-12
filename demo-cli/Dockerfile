FROM golang:1.13-alpine as builder

ENV GO111MODULE on

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk --no-cache add git

WORKDIR /app/policy-demo-cli

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o policy-demo-cli

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache bash ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/policy-demo-cli/policy-demo-cli .

CMD ["./policy-demo-cli"]