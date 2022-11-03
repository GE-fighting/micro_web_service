package server

import (
	"github.com/zsj/micro_web_service/gen/idl/demo"
	"github.com/zsj/micro_web_service/gen/idl/order"
)

type Server struct {
	// 使用unsafe可以强制让编译器检查是否实现了相关方法
	demo.UnsafeDemoServiceServer
	order.UnsafeOrderServiceServer
}
