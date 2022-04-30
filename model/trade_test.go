package model

import (
	"encoding/json"
	"testing"
	"time"
)

func TestSerialization(t *testing.T) {

	serialized, err := json.Marshal(&TradeJ{
		Id:       0,
		DateTime: time.Now(),
		Symbol:   "BTC",
		Price:    1,
		Amount:   2,
	})
	if err != nil {
		return
	}

	println(string(serialized))

}
