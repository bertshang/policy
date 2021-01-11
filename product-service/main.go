package main

import (
	"fmt"
	"github.com/bertshang/policy/product-service/handler"
	"github.com/bertshang/policy/product-service/model"
	"github.com/bertshang/policy/product-service/repo"
	"github.com/bertshang/policy/product-service/db"
	pb "github.com/bertshang/policy/product-service/proto/product"
	"github.com/micro/go-micro"
	wrapperPrometheus "github.com/micro/go-plugins/wrapper/monitoring/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

// 启动 HTTP 服务监听客户端数据采集
func prometheusBoot() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":9092", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

func main()  {
	// 创建数据库连接，程序退出时断开连接
	database, err := db.CreateConnection()
	defer database.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	//数据库迁移（商品、图片、品牌、类目、属性相关数据表）
	database.Set("gorm:table_options", "charset=utf8")
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.ProductImage{})
	database.AutoMigrate(&model.Brand{})
	database.AutoMigrate(&model.Category{})
	database.AutoMigrate(&model.Attribute{})
	database.AutoMigrate(&model.AttributeValue{})
	database.AutoMigrate(&model.ProductAttribute{})

	// 初始化 Repo 实例用于后续数据库操作
	productRepo := &repo.ProductRepository{database}
	imageRepo := &repo.ImageRepository{database}
	brandRepo := &repo.BrandRepository{database}
	categoryRepo := &repo.CategoryRepository{database}
	attributeRepo := &repo.AttributeRepository{database}

	// 以下是 Go Micro 创建微服务流程
	srv := micro.NewService(
		micro.Name("policy.service.product"),
		micro.Version("latest"),  // 新增接口版本参数
		//micro.WrapHandler(wrapperTrace.NewHandlerWrapper()),//trace中间件
		//micro.WrapClient(wrapperTrace.NewClientWrapper()),//trace中间件
		micro.WrapHandler(wrapperPrometheus.NewHandlerWrapper()),//这里是prometheus中间件
	)
	srv.Init()

	// 注册服务处理器
	pb.RegisterProductServiceHandler(srv.Server(), &handler.ProductService{productRepo})
	pb.RegisterImageServiceHandler(srv.Server(), &handler.ImageService{imageRepo})
	pb.RegisterBrandServiceHandler(srv.Server(), &handler.BrandService{brandRepo})
	pb.RegisterCategoryServiceHandler(srv.Server(), &handler.CategoryService{categoryRepo})
	pb.RegisterAttributeServiceHandler(srv.Server(), &handler.AttributeService{attributeRepo})


	// 采集监控数据
	prometheusBoot()
	// 启动商品服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}