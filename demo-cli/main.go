package main

import (
	"context"
	pb "github.com/bertshang/policy/demo-service/proto/demo"
	"github.com/micro/go-micro"
	"log"
)
func main() {
	service := micro.NewService(micro.Name("policy.demo.cli"));
	service.Init();

	client := pb.NewDemoServiceClient("policy.demo.service", service.Client());
	rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "bertshang"});
	if err != nil {
		log.Fatalf("服务调用失败: %v", err);
		return;
	}

	log.Println(rsp.Text);
}
