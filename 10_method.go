package main

import (
	"fmt"
	"strings"
)

/*
  Method adalah fungsi yang menempel pada struct , sehingga hanya bisa di akses lewat
  variabel objek.
  Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga
  level private (level akses nantinya akan dibahas lebih detail pada bab selanjutnya). Dan
  juga, dengan menggunakan method sebuah proses bisa di-enkapsulasi dengan baik.
*/

type pet struct {
	name  string
	owner string
	sound string
}

func (p pet) sayHello() {
	fmt.Println("Halo", p.name)
}

func (p pet) getNameAt(i int) string {
	return strings.Split(p.name, " ")[i-1]
}

func methodBiasa() {
	var p1 = pet{"Lycaa Tjan", "Poppy", "Woof"}
	var name = p1.getNameAt(2)

	p1.sayHello()
	fmt.Println("Nama panggilan: ", name)
}

type petPointer struct {
	name  string
	owner string
	sound string
}

func (p petPointer) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	p.name = name
}

func (p *petPointer) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to ", name)
	p.name = name
}

func methodPointer() {
	var p1 = petPointer{"Lycaa Tjan", "Poppy", "Woof"}

	fmt.Println("p1 before", p1.name)
	p1.changeName1("Miko")
	fmt.Println("p1 after changeName1", p1.name)

	p1.changeName2("Meyer")
	fmt.Println("p1 after changeName2", p1.name)

}

func main() {
	methodBiasa()
	methodPointer()
}
