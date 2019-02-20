package logger

import (
	"log"
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
	log.Printf("📦 %v {\n", t)
	for i := 0; i < t.NumField(); i++ {
		log.Printf("\t%-15s%v\n", t.Field(i).Name, v.Field(i))
	}
	log.Println("}")
}
