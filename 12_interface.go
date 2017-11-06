package main

import (
	"fmt"
	"math"
)

/*
  Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), dan
  dibungkus dengan nama tertentu.
  Interface merupakan tipe data. Nilai objek bertipe interface default-nya adalah nil .
  Interface mulai bisa digunakan jika sudah ada isinya, yaitu objek konkret yang memiliki
  definisi method minimal sama dengan yang ada di interface-nya.
*/

func main() {
	var bangunDatar hitung
	bangunDatar = persegi{10.0}

	fmt.Println("=========Persegi")
	fmt.Println("luas: ", bangunDatar.luas())
	fmt.Println("keliling: ", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}

	fmt.Println("=========lingkaran")
	fmt.Println("luas: ", bangunDatar.luas())
	fmt.Println("keliling: ", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jari())
}

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	d float64
}

type persegi struct {
	sisi float64
}

func (l lingkaran) jari() float64 {
	return l.d / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.d
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}
