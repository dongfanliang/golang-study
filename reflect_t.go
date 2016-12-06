package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintln(u.Age, u.Id, u.Name)
}

func (u User) Hello(m User) (int, string) {
	fmt.Println("Hello", m.Name, ", I'm ", u.Name)
	return m.Age + u.Age, u.Name
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)         //获取接口的类型
	fmt.Println("Type:", t.Name()) //t.Name() 获取接口的名称

	if t.Kind() != reflect.Struct { //通过Kind()来判断反射出的类型是否为需要的类型
		fmt.Println("err: type invalid!")
		return
	}

	v := reflect.ValueOf(o) //获取接口的值类型
	fmt.Printf("Fields: %+v\n", v)

	for i := 0; i < t.NumField(); i++ { //NumField取出这个接口所有的字段数量
		f := t.Field(i)                                   //取得结构体的第i个字段
		val := v.Field(i).Interface()                     //取得字段的值
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val) //第i个字段的名称,类型,值
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m)
		fmt.Printf("%6s: %v\n", m.Name, m.Type) //获取方法的名称和类型
	}
}

func main() {
	u := User{Id: 3, Name: "cat", Age: 34}
	u := User{Id: 3, Name: "cat", Age: 34}
	Info(u)
}
