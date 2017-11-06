package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		11. OPERATOR :
			|| && !
			== <= >=
			% + - * /
		12. CONDITION :
			no ternary
			if else => must use {}, X (), V var temporer
			switch => break udah otomatis, kalau mau lanjut ke case selanjutnya ceritanya ga otomatis break pake fallthrough
	*/
	m := int(time.Now().Month())

	if m%2 == 0 {
		fmt.Println("bulan genap")
	} else {
		fmt.Print("bulan ganjil")
	}

	point := 10040.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if point >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	}

	switch m {
	case 1, 12:
		fmt.Println("Bulan Liburan ya")
	case 7:
		fmt.Println("Bulan Juli")
	case 8:
		fmt.Println("Bulan Agustus")
		fallthrough
	default:
		{
			fmt.Println("Jangan2 kena fallthrough")
			fmt.Println("Bukan Juli / Agustus")
			fmt.Println("Bukan bulan liburan juga loh")
		}
	}
}
