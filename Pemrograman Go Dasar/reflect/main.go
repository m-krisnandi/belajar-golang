package main

import "fmt"
import "reflect"

// akses property variable object
type student struct {
	Name string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama: ", reflectType.Field(i).Name)
		fmt.Println("tipe data: ", reflectType.Field(i).Type)
		fmt.Println("nilai: ", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("type: ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("value: ", reflectValue.Int())
	}

	// pengaksesan nilai dalam bentuk interface{}
	fmt.Println("nilai variable: ", reflectValue.Interface())

	// konversi interface ke tipe aslinya
	fmt.Println("nilai variable: ", reflectValue.Interface().(int))

	fmt.Println("")

	// akses property variable object
	var s1 = &student{Name: "john wick", Grade: 2}
	s1.getPropertyInfo()

	// akses method variable object
	fmt.Println("nama:", s1.Name)

	var reflectValueOfS1 = reflect.ValueOf(s1)
	var method = reflectValueOfS1.MethodByName("SetName")
	method.Call([]reflect.Value {
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama:", s1.Name)
}

