/*
数组是一个结构体
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

数组传参，实际上是这个结构体的副本传递，array所指向的数组没有动；
change_array 修改slice时，只是修改了副本s的array的slice，并不会影响到x的array
change_array_item 修改了s[0],相当于array所指向的数组实体被修改了，所以会影响到外面的x

x := []*MyStruct{&MyStruct{"OldName1"}, &MyStruct{"OldName2"}}
x := []MyStruct{MyStruct{"OldName1"}, MyStruct{"OldName2"}}
效果相同
输出结果：
0xc208039f40
0xc208039f30
0xc208039f30 in change_array start...
0xc208039f38 in change_array end.
OldName1 OldName2
0xc208039f30 in change_array_item start....
0xc208039f30 in change_array_item end.
NewName OldName2

*/
package main

type MyStruct struct {
	Name string
}

func change_array(s []*MyStruct) {
	//println(&s)
	println(&(s[0]), "in change_array start...")
	s = s[1:]
	println(&(s[0]), "in change_array end.")
}

func change_array_item(s []*MyStruct) {
	println(&(s[0]), "in change_array_item start....")
	s[0].Name = "NewName"
	println(&(s[0]), "in change_array_item end.")
}

func doArray() {
	x := []*MyStruct{&MyStruct{"OldName1"}, &MyStruct{"OldName2"}}
	println(&x)      // 数组struct的地址
	println(&(x[0])) // array的地址
	change_array(x)
	println(x[0].Name, x[1].Name)
	change_array_item(x)
	println(x[0].Name, x[1].Name)
}

func main() {
	doArray()
}
