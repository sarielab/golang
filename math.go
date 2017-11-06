package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		mtd    string
		angka1 string
		angka2 string
		result float64
		a      float64
		b      float64
		err    error
		// dialog []string
		// input string
	)

	// dialog = []string{"Pilih metode: [tambah kurang kali bagi]", "Masukan angka1", "Masukan angka2"}
	//
	// for i, d := range dialog {
	// 	fmt.Println(d)
	//     fmt.Scanln(&input)
	// }

	fmt.Println("pilih metode: [tambah kurang kali bagi]")
	fmt.Scanln(&mtd)

	fmt.Println("Masukan angka1")
	fmt.Scanln(&angka1)

	a, err = strconv.ParseFloat(angka1, 64)
	if err != nil {
		fmt.Println("input harus berupa angka")
	}

	fmt.Println("Masukan angka2")
	fmt.Scanln(&angka2)
	b, err = strconv.ParseFloat(angka2, 64)
	if err != nil {
		fmt.Println("input harus berupa angka")
	}

	fmt.Println("angka", a, b)

	switch mtd {
	case "tambah":
		result = a + b
	case "kurang":
		result = a - b
	case "bagi":
		result = a / b
	case "kali":
		result = a * b
	default:
		result = 0
	}

	fmt.Println("result", result)
}
