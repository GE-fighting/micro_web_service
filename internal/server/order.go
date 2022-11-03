package server

import (
	"context"
	"errors"
	"time"

	"github.com/zsj/micro_web_service/gen/idl/order"
	"github.com/zsj/micro_web_service/internal/model"
	"github.com/zsj/micro_web_service/internal/service"
	"github.com/zsj/micro_web_service/internal/zlog"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) ListOrders(ctx context.Context, request *order.ListOrdersRequest) (*order.ListOrdersResponse, error) {
	orders, err := service.NewOrderService().List(ctx, int(request.PageNum), int(request.PageSize), nil)
	if err != nil {
		zlog.Suagr.Fatalf("call service.List() failed,%v", err)
		return nil, err
	}
	// resp := new(order.ListOrdersResponse)
	resultsOrders := make([]*order.Order, len(orders))
	for k, v := range orders {
		// resultsOrders[k].Id = int32(orders[k].ID)
		// resultsOrders[k].Name = orders[k].Name
		// resultsOrders[k].Price = orders[k].Price
		// resultsOrders[k].CreateTime = timestamppb.New(orders[k].CreateTime)
		resultsOrders[k] = &order.Order{
			Id:         int32(v.ID),
			Name:       v.Name,
			Price:      v.Price,
			CreateTime: timestamppb.New(v.CreateTime),
		}
	}
	return &order.ListOrdersResponse{Orders: resultsOrders, Count: int32(len(orders))}, nil
}

func (s *Server) CreateOrder(ctx context.Context, request *order.CreateOrderRequest) (*order.Order, error) {
	order := &model.Order{
		Name:       request.Order.Name,
		Price:      request.Order.Price,
		CreateTime: time.Now(),
	}
	err := service.NewOrderService().Add(ctx, order)
	if err != nil {
		zlog.Suagr.Fatalf("create order failed : %v", err)
		return nil, err
	}
	request.Order.Id = int32(order.ID)
	return request.Order, nil
}

func (s *Server) UpdateOrder(ctx context.Context, resp *order.UpdateOrderRequest) (*emptypb.Empty, error) {
	orderUpdate := &model.OrderFields{
		Order: &model.Order{
			Name:  resp.Order.Name,
			Price: resp.Order.Price,
		},
		Fields: resp.UpdateMask.Paths,
	}
	orderCondition := &model.OrderFields{
		Order: &model.Order{
			ID: int64(resp.Order.Id),
		},
		Fields: []string{"id"},
	}
	err := service.NewOrderService().Update(ctx, orderUpdate, orderCondition)
	if err != nil {
		zlog.Suagr.Fatalf("update order failed,%v \n", err)
	}
	return &emptypb.Empty{}, err
}

func (s *Server) GetOrder(ctx context.Context, resp *order.GetOrderRequest) (*order.Order, error) {
	orderQuery := &model.OrderFields{
		Order:  &model.Order{Name: resp.Name},
		Fields: []string{"name"},
	}
	orders, err := service.NewOrderService().List(ctx, 0, 1, orderQuery)
	if err != nil {
		zlog.Suagr.Fatalf("get order failed,%v \n", err)
		return nil, err
	} else if len(orders) == 0 {
		return nil, errors.New("no match result")
	}
	order := &order.Order{
		Id:         int32(orders[0].ID),
		Name:       orders[0].Name,
		Price:      orders[0].Price,
		CreateTime: timestamppb.New(orders[0].CreateTime),
	}
	return order, nil
}

func (s *Server) DeleteOrder(ctx context.Context, req *order.DeleteOrderRequest) (*emptypb.Empty, error) {
	condition := &model.OrderFields{
		Order:  &model.Order{Name: req.Name},
		Fields: []string{"name"},
	}

	// TODO soft delete
	updated := &model.OrderFields{
		Order:  &model.Order{},
		Fields: []string{},
	}

	return &emptypb.Empty{}, service.NewOrderService().Update(ctx, updated, condition)
}
