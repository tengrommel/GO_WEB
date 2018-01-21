package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	Id int	`json:"id"`
}

func (s Student)String() string {
	return fmt.Sprintf("name: %s, id: %d", s.Name, s.Id)
}

func print_any(x interface{})  {
	t := reflect.TypeOf(x)
	fmt.Println(t.Kind())
	t = t.Elem()
	fmt.Println(t.Name())
	fmt.Println(t.PkgPath())
	for i:=0;i<t.NumField();i++{
		field := t.Field(i)
		fmt.Println(field)
		fmt.Println("jsonkey: ", field.Tag.Get("json"))
	}
	field, _ := t.FieldByName("Name")
	fmt.Printf("%#v\n", field)
	for i:=0;i<t.NumMethod();i++{
		method := t.Method(i)
		fmt.Println(method.Name)
	}

	v := reflect.ValueOf(x).Elem()
	vfiled := v.FieldByName("Name")
	fmt.Println(vfiled.String())

	method := v.MethodByName("String")
	ret := method.Call(nil)
	fmt.Println(ret[0].String())
}

func main() {
	s := &Student{
		Name:"teng",
		Id: 1,
	}
	print_any(s)
}
