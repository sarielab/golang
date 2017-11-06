package main

/*
  import
    1. . bikin selevel
      b4 var s1 = library.Student{"Ernie", 21}
      now var s1 = Student{"Ernie", 21}
    2. f "fmt" --> aliasing

  Jalankan partial package => 1 pack beda file
    run sekaligus
*/

import (
	"fmt"
	. "latihan/11_accessLevel/library"
)

func main() {
	sayHello("Ethan")
	// test.SayTest()
	// packageStruct()

	fmt.Printf("Name: %s\n", Student.Name)
	fmt.Printf("Grade: %d\n", Student.Grade)
}

// func packageStruct() {
// var s1 = Student{"Ernie", 21}
// f.Println("name", s1.Name)
// f.Println("grade", s1.Grade)
// }
