package main

import (
	"fmt"

	"github.com/MasterOfBinary/go-boilerplate/hello"
)

func notUsed() {
	fmt.Println("nooo")
}

func main() {
	str := hello.Hello()
	fmt.Println(str)
}
