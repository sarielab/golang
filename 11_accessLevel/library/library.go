package library

/*
  ACCESS MODIFIER dari package lain
  - name?
    capitalize => public
    small => private
	- apa aja? struct, method, props param
	- step
    + nama package = nama folder
    + introduce => private
    + SayHello => public
*/

import "fmt"

// type Student struct {
// 	Name  string
// 	Grade int
// }

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "Lithany"
	Student.Grade = 2

	fmt.Println("Importing Library")
}

func sayHello() {
	fmt.Println("hello")
}

func introduce(name string) {
	fmt.Println("nama saya", name)
}
