package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/zsj/micro_web_service/gen/idl/demo"
	"github.com/zsj/micro_web_service/gen/idl/order"
	"github.com/zsj/micro_web_service/internal/config"
	"github.com/zsj/micro_web_service/internal/mysql"
	"github.com/zsj/micro_web_service/internal/server"
	"github.com/zsj/micro_web_service/internal/zlog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", ":9090", "gRPC server endpoint")
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := demo.RegisterDemoServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return errors.Wrap(err, "RegisterDemoServiceHandlerFromEndpoint error")
	} else if err := order.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", config.Viper.GetInt("server.grpc.port")), opts); err != nil {
		return errors.Wrap(err, "RegisterOrderServiceHandlerFromEndpoint error")
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Viper.GetInt("server.http.port")), mux)
}

func main() {
	configPath := flag.String("c", "./", "config file path")
	flag.Parse()

	if err := config.Load(*configPath); err != nil {
		panic(err)
	}
	zlog.Init(config.Viper.GetString("zlog.path"))
	defer zlog.Sync()
	// 初始化mysql
	if err := mysql.Init(config.Viper.GetString("mysql.user"), config.Viper.GetString("mysql.password"), config.Viper.GetString("mysql.ipaddress"), config.Viper.GetInt("mysql.port"),
		config.Viper.GetString("mysql.dbName")); err != nil {
		zlog.Suagr.Fatalf("init mysql fail %v", err)
	}
	// 模型迁移到数据库，创建表
	// if err := dao.Migrate(); err != nil {
	// 	zlog.Suagr.Fatalf("migrate faid %v", err)
	// }
	go func() {
		port := config.Viper.GetInt("server.grpc.port")
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			panic(err)
		}

		s := grpc.NewServer()
		demo.RegisterDemoServiceServer(s, &server.Server{})
		order.RegisterOrderServiceServer(s, &server.Server{})
		if err = s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	if err := run(); err != nil {
		panic(err)
	}
}
