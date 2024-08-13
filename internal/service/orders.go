package service

import (
	"AsiaYo-Backend-pre-test/internal/model"
)

// CreateOrder 建立訂單
func CreateOrder(in *model.CreateOrder) error {
	// 格式檢查
	if err := in.Validate(); err != nil {
		return err
	}

	// 進行資料庫操作
	// ...

	return nil
}
