package Log

import (
	"fmt"
	"reflect"
)

type Logger interface {
	Log()
}

func PrintFields(target Logger) {
	v := reflect.Indirect(reflect.ValueOf(target))
	t := v.Type()
	fmt.Println("{")
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("\t%-15s%v\n", t.Field(i).Name, v.Field(i))
	}
	fmt.Println("}")
}
