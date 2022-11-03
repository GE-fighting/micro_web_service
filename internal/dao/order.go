package dao

import (
	"errors"

	"github.com/zsj/micro_web_service/internal/model"
	"github.com/zsj/micro_web_service/internal/mysql"
	"gorm.io/gorm"
)

// type Order struct {
// 	ID         int64
// 	Name       string    `gorm:"comment:订单名"`
// 	Price      float32   `gorm:"comment:订单价格"`
// 	CreateTime time.Time `gorm:"comment:订单创建时间"`
// }

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{db: mysql.GormDB}
}

// func NewOrderFields(order *model.Order, fields []interface{}) *model.OrderFields {
// 	return &model.OrderFields{
// 		order:  *order,
// 		fields: fields,
// 	}
// }

func (orderRepo *OrderRepo) AddOrder(order *model.Order) (err error) {
	err = orderRepo.db.Create(order).Error
	return
}

func (repo *OrderRepo) QueryOrders(pageNum, pageSie int, condition *model.OrderFields) (orders []model.Order, err error) {
	db := repo.db
	if condition != nil {
		db.Where(condition.Order, condition.Fields)
	}
	err = db.Limit(pageSie).Offset((pageNum - 1) * pageSie).Find(&orders).Error
	return
}

func (repo *OrderRepo) UpdateOrder(updated, condition *model.OrderFields) (err error) {
	if updated == nil || len(updated.Fields) == 0 {
		return errors.New("update must choose certail fields")
	} else if condition == nil {
		return errors.New("update must include where condition")
	}
	err = repo.db.Select(updated.Fields).Where(condition.Order, condition.Fields).Updates(updated.Order).Error
	return
}

// 迁移模型，创建表
func Migrate() error {
	return mysql.GormDB.AutoMigrate(&model.Order{})
}

// 创建一条订单记录
// func CreateOrder(order *Order) (err error) {
// 	err = mysql.GormDB.Create(order).Error
// 	return
// }

// 查询符号条件的订单记录
// func QueryOrders(pageNumber, pageSize int, condition *Order, fields ...interface{}) (orders []Order, err error) {
// 	err = mysql.GormDB.Where(condition, fields...).Limit(pageNumber).
// 		Offset((pageNumber - 1) * pageSize).Find(&orders).Error
// 	return
// }

// 更新某条订单记录
// func UpdateOrder(order *Order, fileds ...interface{}) (err error) {
// 	err = mysql.GormDB.Model(&order).Select(fileds).Updates(order).Error
// 	return
// }
