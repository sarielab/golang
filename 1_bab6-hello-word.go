package main

import (
	"fmt"
)

// harus ada muncul dan dijalankan pertamax
func main() {
	// var firstName string = "Poppy"
	var exist bool = true
	const firstName string = "John"
	const lastNameCONST = "DOE"

	message := `Saya "manusia".
	Enter dimana mana.
	Biar cobain aja.`
	lastName := "Sari"
	decimalNumber := 2.62
	// _ = "tess"

	// fmt.Print(_)
	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"
	//
	// var fourth, fifth, sixth string = "empat", "lima", "enam"
	//
	// seventh, eight, ninth := "tujuh", 8, true

	fmt.Printf("halo %s %s! %s \n", firstName, lastName, lastNameCONST)
	fmt.Println("hello", "how are you", "<isi-pesan>")
	fmt.Printf("desimal: %f\n", decimalNumber)
	fmt.Printf("desimal: %.3f\n", decimalNumber)
	fmt.Printf("boolean: %t\n", exist)
	fmt.Println(message)

}
