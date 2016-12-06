package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("failure", "urei & i"))
	fmt.Println(strings.Count("five", "iv")) // before & after each rune
	fmt.Println(strings.Fields(" five	"))
	fmt.Printf("%qxxx\n", strings.Split(" xyz ", " "))
}
