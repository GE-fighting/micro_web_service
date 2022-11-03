package service

import (
	"context"

	"github.com/zsj/micro_web_service/internal/dao"
	"github.com/zsj/micro_web_service/internal/model"
	"github.com/zsj/micro_web_service/internal/zlog"
)

type OrderService struct {
	OrderRepo model.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		OrderRepo: dao.NewOrderRepo(),
	}
}

func (orderService *OrderService) List(ctx context.Context, pageNum, pageSize int, condition *model.OrderFields) (orders []model.Order, err error) {
	orders, err = orderService.OrderRepo.QueryOrders(pageNum, pageSize, condition)
	if err != nil {
		zlog.Suagr.Fatalf("Orderservice lsit failed : %v", err)
		return nil, err
	}
	return orders, nil
}

func (orderService *OrderService) Add(ctx context.Context, order *model.Order) error {
	err := orderService.OrderRepo.AddOrder(order)
	if err != nil {
		zlog.Suagr.Fatalf("OrderService add Order failed : %v", err)
	}
	return err
}

func (orderService *OrderService) Update(ctx context.Context, updated, condition *model.OrderFields) error {
	err := orderService.OrderRepo.UpdateOrder(updated, condition)
	if err != nil {
		zlog.Suagr.Fatalf("OrderService update Order failed : %v", err)
	}
	return err
}
