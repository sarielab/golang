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

type student struct {
	name  string
	grade int
}

func (st student) getNameAt(i int) string {
	return strings.Split(st.name, " ")[i-1]
}

func (st student) hello() {
	fmt.Println("halo", st.name)
}

func method1() {

	var s1 = student{"Lycaa D", 17}
	var name = s1.getNameAt(2)

	s1.hello()
	fmt.Println("nama panggilan: ", name)
}

func main() {
	method1()
}
