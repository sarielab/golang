package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = flag.String("name", "anon", "type your name")
	var age = flag.Int64("age", 65, "type your age")

	flag.Parse()
	fmt.Printf("name\t:%s\n", *name)
	fmt.Printf("age\t:%d\n", *age)
}

//go run 26_flag.go -name="john wick" -age=22
