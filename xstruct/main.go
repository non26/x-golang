package main

import (
	"fmt"
	"reflect"
)

func main() {
	type employeeDetails struct {
		id          int16
		name        string
		designation string
	}

	x := employeeDetails{
		id:          1,
		name:        "N",
		designation: "N",
	}

	// fields := reflect.VisibleFields(reflect.TypeOf(x))
	// for _, field := range fields {
	// 	fmt.Printf("Key: %s\tType: %s\n", field.Name, field.Type)
	// }

	// v := reflect.ValueOf(x)
	// for i := 0; i < v.NumField(); i++ {
	// 	fmt.Println(v.Type().Field(i).Name)
	// 	fmt.Println("\t", v.Field(i))
	// }
	sign(x)
}

func sign(m interface{}) {
	v := reflect.ValueOf(m)
	var s string
	for i := 0; i < v.NumField(); i++ {
		s = s + fmt.Sprintf("%v=%v", v.Type().Field(i).Name, v.Field(i))
		if i != v.NumField()-1 {
			s = s + "&"
		}
	}

	println(s)
}
