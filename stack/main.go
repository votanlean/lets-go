package main

import (
	"fmt"

	"stack/mystack"
)

var st1 mystack.Stack
func main() {
	st1.Push(1)	
	st1.Push("hi")
	st1.Push(3.5)
	st1.Push([]string{"Java", "C++"})
	for {
		item, err := st1.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}

}