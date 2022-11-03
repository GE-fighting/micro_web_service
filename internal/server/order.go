package server

import (
	"context"

	"github.com/zsj/micro_web_service/gen/idl/order"
	"github.com/zsj/micro_web_service/internal/service"
	"github.com/zsj/micro_web_service/internal/zlog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) ListOrders(ctx context.Context, request *order.ListOrdersRequest) (*order.ListOrdersResponse, error) {
	orders, err := service.NewOrderService().List(ctx, int(request.PageNum), int(request.PageSize), nil)
	if err != nil {
		zlog.Suagr.Fatalf("call service.List() failed,%v", err)
		return nil, err
	}
	resultsOrders := make([]*order.Order, len(orders))
	for k, _ := range resultsOrders {
		resultsOrders[k].Id = int32(orders[k].ID)
		resultsOrders[k].Name = orders[k].Name
		resultsOrders[k].Price = orders[k].Price
		resultsOrders[k].CreateTime = timestamppb.New(orders[k].CreateTime)
	}
	return &order.ListOrdersResponse{Orders: resultsOrders, Count: int32(len(orders))}, nil
}
