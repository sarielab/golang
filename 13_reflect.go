package main

/*
  reflect utk dapatkan info ttg variable
  fungsi:
    - reflect.ValueOf()
    - reflect.TypeOf()
  jika mau loopStruct utk cek
    var reflectValue = reflect.ValueOf(s)
  	if reflectValue.Kind() == reflect.Ptr { //pointer bukan?
  		reflectValue = reflectValue.Elem()
  	}

  	var reflectType = reflectValue.Type()

  	for i := 0; i < reflectValue.NumField(); i++ {
  		fmt.Println("Name : ", reflectType.Field(i).Name)
  		fmt.Println("Type : ", reflectType.Field(i).Type)
  		fmt.Println("Grrade : ", reflectValue.Field(i).Interface()) // hasil apapun oke
  	}

  reflectValue = reflect.ValueOf(s)
    .Interface() => terima semua
    .Ptr => pointer
    .Kind() => tipedata


  method.Call() => exec method

*/

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func (s *student) getStructPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	if reflectValue.Kind() == reflect.Ptr { //pointer bukan?
		reflectValue = reflectValue.Elem()
	}

	fmt.Println("reflectValue elem", reflectValue)

	var reflectType = reflectValue.Type()
	fmt.Println("\nreflectType", reflectType)

	fmt.Println("\n\n-------------------------\n")
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("Name : ", reflectType.Field(i).Name)
		fmt.Println("Type : ", reflectType.Field(i).Type)
		fmt.Println("Grrade : ", reflectValue.Field(i).Interface())
	}
	fmt.Println("reflectValue elem", reflectValue)
}

func struct1() {
	var number int = 77
	var reflectValue = reflect.ValueOf(number)
	var reflectType = reflect.TypeOf(number)
	fmt.Println("valueof ", reflectValue)
	fmt.Println("typeof ", reflectType)
}

func (s *student) getMethodStructProperty() {
	var reflectValue = reflect.ValueOf(s)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Wick"),
	})
	fmt.Println("nama", s.Name)
}

func main() {
	struct1()
	fmt.Println("\n\n")
	var s1 = &student{Name: "Wick", Grade: 2}
	s1.getStructPropertyInfo()
	fmt.Println("\n\n")
	s1.getMethodStructProperty()
}
