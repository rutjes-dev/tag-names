package main

import (
	"tag-names/marshalling"
	"tag-names/model"
	"testing"
	"time"
)

func BenchmarkReflectionMarshal(b *testing.B) {
	tr := model.Trade{
		Id:       1,
		DateTime: time.Now(),
		Symbol:   "BTC",
		Price:    60000.00,
		Amount:   1.0,
	}

	for i := 0; i < b.N; i++ {

		MarshalReflection(tr)

	}

}

func BenchmarkHardcodedMarshal(b *testing.B) {
	tr := model.Trade{
		Id:       1,
		DateTime: time.Now(),
		Symbol:   "BTC",
		Price:    60000.00,
		Amount:   1.0,
	}

	for i := 0; i < b.N; i++ {
		MarshalHardcoded(tr)
	}
}

func BenchmarkGeneratedMarshal(b *testing.B) {
	tr := model.Trade{
		Id:       1,
		DateTime: time.Now(),
		Symbol:   "BTC",
		Price:    60000.00,
	}

	for i := 0; i < b.N; i++ {
		marshalling.MarshalTrade(tr)
	}
}
