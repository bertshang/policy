package main

import (
	"context"
	"fmt"
	"github.com/bertshang/policy/common/tracer"
	pb "github.com/bertshang/policy/demo-service/proto/demo"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/bertshang/policy/common/wrapper/breaker/hystrix"
	traceplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	hystrix.Configure([]string{"policy.service.demo.DemoService.SayHello"})
	// 初始化追踪器
	t, io, err := tracer.NewTracer("policy.demo.cli", os.Getenv("MICRO_TRACE_SERVER"))
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()



	service := micro.NewService(
		micro.Name("policy.demo.cli"),
		micro.WrapClient(traceplugin.NewClientWrapper(t)),
		micro.WrapClient(hystrix.NewClientWrapper()),
	)
	service.Init()

	client := pb.NewDemoServiceClient("policy.service.demo", service.Client())

	// 创建空的上下文, 生成追踪 span
	span, ctx := opentracing.StartSpanFromContext(context.Background(), "call")
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}
	defer span.Finish()

	// 注入 opentracing textmap 到空的上下文用于追踪
	opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))
	ctx = opentracing.ContextWithSpan(ctx, span)
	ctx = metadata.NewContext(ctx, md)
	// 记录请求 && 响应 && 错误
	req := &pb.DemoRequest{Name: "bertshang"}
	span.SetTag("req", req)
	resp, err := client.SayHello(ctx, req)
	if err != nil {
		span.SetTag("err", err)
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	span.SetTag("resp", resp)


	// 模拟常驻内存
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case os.Interrupt, os.Kill, syscall.SIGQUIT:
				fmt.Println("退出客户端")
				os.Exit(0)
			default:
				fmt.Println("程序执行中...")
			}
		}
	}()

	for {
		rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "bertshang"})
		if err != nil {
			log.Fatalf("服务调用失败：%v", err)
			return
		}
		log.Println(rsp.Text)
		time.Sleep(3 * time.Second)
	}
	log.Println(resp.Text)
}