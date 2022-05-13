package order

import "time"

type GetOrderListParams struct {
	Status            string
	StartOrderDate    time.Time
	EndOrderDate      time.Time
	SortDateAscending bool
	Limit             int64
	Page              int64
}

type OrderProductDetail struct {
	ProductName      string
	Quantity         uint32
	PricePerItem_IDR uint32
}

type OrderList struct {
	OrderID      string    `bson:"_id"`
	BusinessName string    `bson:"business_name"`
	OrderTime    time.Time `bson:"order_time"`
	TotalPayIDR  uint32    `bson:"total_pay_idr"`
	Status       string    `bson:"status"`
}
