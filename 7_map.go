package main

import (
	"fmt"
)

/*
Map adalah tipe data asosiatif yang ada di Golang, berbentuk key-value.
Untuk setiap data (atau value) yang disimpan, disiapkan juga key-nya.
Key harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang bersangkutan.
Kalau dilihat, map mirip seperti slice,
hanya saja indeks yang digunakan untuk pengaksesan bisa ditentukan sendiri tipe-nya (indeks tersebut adalah key).
*/

/*
  MAP
  - deklarasi
    1. tanpa nilai awal
      var zodiac = map[string]int{}
      var zodiac = make(map[string]int)
      var zodiac = *new(map[string]int)
    2. vertikal => var zodiac = map[string]int{"aries":1, "taurus":2}
    3. horizontal => var zodiac = map[string]int{
      "aries" : 1,
      "taurus": 2,
    }
  - delete => delete(zodiac, "aries")
  - isExist => var value, isExist = zodiac["aries"]
  - combine slice and map => array of objects
    var pets = []map[string]string{
      {"name": "lycaa", "hobby": "playing", "gender": "female"},
    }
*/

var zodiac = map[string]int{
	"aries":   1,
	"taurus":  2,
	"cancer":  4,
	"scorpio": 7,
}

func main() {
	// loopRange()
	// deleteItem()
	// searchItem()
	sliceMap()
}

func loopRange() {
	for key, val := range zodiac {
		fmt.Printf("zodiac number %s is %d \n", key, val)
	}
}

func deleteItem() {
	fmt.Println("=============DELETE==============")
	fmt.Println(len(zodiac))
	fmt.Println(zodiac)

	delete(zodiac, "aries")

	fmt.Println(len(zodiac))
	fmt.Println(zodiac)
}

func searchItem() {
	var val, isExist = zodiac["virgo"]

	fmt.Println("==============SEARCH=============")
	if isExist {
		fmt.Println(val)
	} else {
		fmt.Println("virgo is not exist")
	}
}

func sliceMap() {
	var animals = []map[string]string{
		{"name": "Lycaa", "gender": "female", "hobby": "sleep"},
		{"name": "Miko", "gender": "male"},
		{"name": "Meyer", "gender": "male"},
	}

	for _, animal := range animals {
		fmt.Println(animal["name"], animal["gender"])
	}

}
