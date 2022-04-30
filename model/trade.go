package model

import "time"

type Trade struct {
	Id       int64     `map:"id"`
	DateTime time.Time `map:"date_time"`
	Symbol   string    `map:"symbol"`
	Price    float64   `map:"price"`
	Amount   float64   `map:"amount"`
}

type TradeJ struct {
	Id       int64     `json:"id" map:"id"`
	DateTime time.Time `json:"date_time"`
	Symbol   string    `json:"symbol,omitempty"`

	Price float64 `json:"price" validate:"gte=0,lte=130"`

	Amount float64 `json:"amount"`
}

type TradeS struct {
	Id       int64
	DateTime time.Time
	Symbol   string
	Price    float64
	Amount   float64
}
