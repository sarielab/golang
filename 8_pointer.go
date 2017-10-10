package main

import (
	"fmt"
)

/*
  Pointer adalah reference atau alamat memory. Variabel pointer berarti variabel yang berisi
  alamat memori suatu nilai. Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4,
  maka yang dimaksud pointer adalah alamat memori dimana nilai 4 disimpan, bukan nilai
  4 nya sendiri.
  Variabel-variabel yang memiliki reference atau alamat pointer yang sama, saling
  berhubungan satu sama lain dan nilainya pasti sama. Ketika ada perubahan nilai, maka
  akan memberikan efek kepada variabel lain (yang referensi-nya sama) yaitu nilainya ikut
  berubah
*/

func main() {
	varPointer()
	parameterPointer()
}

func varPointer() {
	var numberA int = 6
	var numberB *int = &numberA

	fmt.Println("numberA (value) : ", numberA)
	fmt.Println("numberA (address) : ", &numberA)
	fmt.Println("numberB (value) : ", *numberB)
	fmt.Println("numberB (address) : ", numberB)
	fmt.Println("")

	numberA = 5
	fmt.Println("numberA = 5")
	fmt.Println("numberA (value) : ", numberA)
	fmt.Println("numberA (address) : ", &numberA)
	fmt.Println("numberB (value) : ", *numberB)
	fmt.Println("numberB (address) : ", numberB)
}

func parameterPointer() {
	var change = func(original *int, value int) {
		*original = value
	}

	var number = 4
	fmt.Println("before: ", number) //4

	change(&number, 8)
	fmt.Println("after: ", number) //8

}
