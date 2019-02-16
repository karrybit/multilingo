package log

import (
	"fmt"
	"reflect"
)

// Logger -
type Logger interface {
	Log()
}

// PrintFields is standard output of all properties
func PrintFields(target Logger) {
	v := reflect.Indirect(reflect.ValueOf(target))
	t := v.Type()
	fmt.Printf("📦 %v {\n", t)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("\t%-15s%v\n", t.Field(i).Name, v.Field(i))
	}
	fmt.Println("}")
}
