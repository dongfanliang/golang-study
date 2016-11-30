package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
)

type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
	Sex  int
}

func (this *Person) String() string {
	buffer := bytes.NewBufferString("This is ")
	buffer.WriteString(this.Name + ", ")
	if this.Sex == 0 {
		buffer.WriteString("He ")
	} else {
		buffer.WriteString("She ")
	}
	buffer.WriteString("is ")
	buffer.WriteString(strconv.Itoa(this.Age))
	buffer.WriteString(" years old.")
	return buffer.String()
}

func main() {
	t := &Person{"hell", 34, 0}
	print(t, "\n")
	fmt.Println(t)
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#+v\n", t)
	fmt.Printf("%%\n")
	// 可用字符串拼接
	result1 := fmt.Sprintln("studygolang.com", 2013)
	result2 := fmt.Sprint("studygolang.com", 2013)
	log.Println(result1)
	log.Println(result2)

}
