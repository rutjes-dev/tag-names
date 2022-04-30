package main

import (
	"bytes"
	"fmt"
	"reflect"
	"tag-names/marshalling"
	"tag-names/model"
	"time"
)

const tagName = "map"

func MarshalReflection(st any) string {

	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	rv := reflect.ValueOf(st)

	// TypeOf returns the reflection Type that represents the dynamic type of i.
	// If i is a nil interface value, TypeOf returns nil.
	rt := reflect.TypeOf(st)

	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	rv = reflect.ValueOf(st)

	buf := new(bytes.Buffer)

	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	for i := 0; i < rt.NumField(); i++ {
		tf := rt.Field(i)
		vf := rv.Field(i)
		tag := tf.Tag.Get(tagName)

		// skip nil value
		if vf.IsZero() {
			continue
		}
		switch vf.Interface().(type) {
		case int64:
			buf.WriteString(fmt.Sprintf("%s:%d\n", tag, vf.Int()))
		case time.Time:
			buf.WriteString(fmt.Sprintf("%s:%s\n", tag,
				vf.Interface().(time.Time).Format(time.RFC3339)))
		case string:
			buf.WriteString(fmt.Sprintf("%s:%s\n", tag, vf.String()))
		case float64:
			buf.WriteString(fmt.Sprintf("%s:%.2f\n", tag, vf.Float()))
		}
	}

	return buf.String()
}

func MarshalHardcoded(object model.Trade) string {
	return fmt.Sprintf("id:%d\ndate_time:%s\nsymbol:%s\nprice:%.2f\namount:%.2f\n",
		object.Id,
		object.DateTime.Format(time.RFC3339),
		object.Symbol,
		object.Price,
		object.Amount)

}

func main() {
	tr := model.Trade{
		Id:       1,
		DateTime: time.Now(),
		Symbol:   "BTC",
		Price:    60000.00,
		Amount:   1.0,
	}

	println("Marshal hardcoded by hand")
	println(MarshalHardcoded(tr))
	println("Marshal by reflection")
	println(MarshalReflection(tr))
	println("Marshal by generated code")

	println(marshalling.MarshalTrade(tr))
}
