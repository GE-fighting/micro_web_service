package dao

import (
	"time"

	"github.com/zsj/micro_web_service/internal/mysql"
)

type Order struct {
	ID         int64
	Name       string    `gorm:"comment:订单名"`
	Price      float32   `gorm:"comment:订单价格"`
	CreateTime time.Time `gorm:"comment:订单创建时间"`
}

// 迁移模型，创建表
func Migrate() error {
	return mysql.GormDB.AutoMigrate(&Order{})
}

// 创建一条订单记录
func CreateOrder(order *Order) (err error) {
	err = mysql.GormDB.Create(order).Error
	return
}

// 查询符号条件的订单记录
func queryOrders(pageNumber, pageSize int, condition *Order, fields ...interface{}) (orders []Order, err error) {
	err = mysql.GormDB.Where(condition, fields...).Limit(pageNumber).
		Offset((pageNumber - 1) * pageSize).Find(&orders).Error
	return
}

// 更新某条订单记录
func updateOrder(order *Order, fileds ...interface{}) (err error) {
	err = mysql.GormDB.Model(&order).Select(fileds).Updates(order).Error
	return
}
