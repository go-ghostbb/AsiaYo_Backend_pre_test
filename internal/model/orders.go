package model

import (
	"AsiaYo-Backend-pre-test/internal/utils/stringx"
	"errors"
	"strconv"
)

var (
	ErrNameContainNonEnglishChar = errors.New("Name contains non-English characters")
	ErrNameIsNotCapitalized      = errors.New("Name is not capitalized")
	ErrPriceIsOver2000           = errors.New("Price is over 2000")
	ErrCurrencyFormatWrong       = errors.New("Currency format is wrong")
)

type CreateOrder struct {
	ID       string  `json:"id" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Address  Address `json:"address" binding:"required"`
	Price    string  `json:"price" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

type Address struct {
	City     string `json:"city" binding:"required"`
	District string `json:"district" binding:"required"`
	Street   string `json:"street" binding:"required"`
}

func (c *CreateOrder) Validate() error {
	// 判斷是否包含非英文字母(name)
	if stringx.ContainNonEnglish(c.Name) {
		return ErrNameContainNonEnglishChar
	}

	// 判斷每個單字首字母是否為大寫(name)
	if !stringx.IsFirstLetterUp(c.Name) {
		return ErrNameIsNotCapitalized
	}

	// 判斷貨幣格式是否為TWD或USD
	if !(c.Currency == "TWD" || c.Currency == "USD") {
		return ErrCurrencyFormatWrong
	}

	// string convert to int(price)
	var (
		price int
		err   error
	)
	price, err = strconv.Atoi(c.Price)

	// 如果貨幣格式為USD, 將USD改為TWD
	if c.Currency == "USD" {
		// 乘上匯率(*31)
		price *= 31
		c.Price = strconv.Itoa(price)
		// 將USD改為TWD
		c.Currency = "TWD"
	}

	// 訂單金額超過2000
	if err != nil {
		return err
	} else if price > 2000 {
		return ErrPriceIsOver2000
	}

	return nil
}
