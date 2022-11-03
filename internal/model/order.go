package model

import (
	"time"
)

type Order struct {
	ID         int64
	Name       string
	Price      float32
	CreateTime time.Time
}

type OrderFields struct {
	Order  *Order
	Fields []string
}

type OrderRepository interface {
	AddOrder(order *Order) (err error)
	QueryOrders(pageNum, pageSie int, condition *OrderFields) (orders []Order, err error)
	UpdateOrder(updated, condition *OrderFields) (err error)
}
