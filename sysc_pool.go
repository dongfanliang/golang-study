package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

func main() {
	b := bufferPool.Get().(*bytes.Buffer)
	b.Reset()
	defer bufferPool.Put(b)
	b.WriteString("hello World")
	fmt.Println(b.String())
}
