package main

import (
	"log"
)

type P struct {
	Name string
}

func work(x []*P) []*P {
	x[0].Name = "love"
	println(&x, &(x[0]))
	return x
}

func main() {
	x := []*P{&P{"hello"}, &P{"world"}}
	println(&x, &x[0])

	a := work(x)
	println(&(a), &(a[0]))
	a = nil
	a = []*P{&P{"hello"}}
	println(&(a), &(a[0]))

	log.Println(x[0].Name)
}
