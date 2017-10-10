package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

var names = []string{"John", "Wick"}
var numbers = []int{2, 10, 11, 8, 9, 0, 8, 3, 1}

var getMinMax = func(numbers []int) (int, int) {
	var min, max int
	min = numbers[0]
	max = numbers[0]

	for _, number := range numbers {
		switch {
		case number > max:
			max = number
		case number < min:
			min = number
		}
	}
	return min, max
}

/*
  rand.Seed(time.Now().Unix())
  memastikan bahwa angka random yang akan di-generate benar-benar acak. Kita bisa gunakan angka apa saja sebagai nilai parameter fungsi ini (umumnya diisi time.Now().Unix()).

  FUNGSI
  - return value
    func nama(params tipeparam) tipereturn { return value}
  - declare param tipe data sama
    func nama(param1, param2, param3 tipe) tipereturn
  - declare param beda beda
    func nama(param1 tipe1, param2 tipe2)
	- multiple return value
		func nama(p1, p2 tipe) (tipe1, tipe2) { return v1, v2}}
	- predifened return value
	- variadic
		var numbers = []int{2,4,5,61,1}
		var avg = calculate(numbers...)
		func avg(numbers ...int) {}
			Deklarasi parameter variadic sama dengan cara deklarasi variabel biasa, pembedanya
			adalah pada parameter jenis ini ditambahkan tanda titik tiga kali ( ... ) tepat setelah
			penulisan variabel (sebelum tipe data). Nantinya semua nilai yang disisipkan sebagai
			parameter akan ditampung oleh variabel tersebut.
	- closure
		func tanpa nama disimpan di variabel
	- IIFE (immediately-invoked function expression)
	- casting
		fmt.Sprintf()
		float64()
*/

func main() {

	printMessage("halo", names)
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	//closure => balikin nilai
	// min, max := getMinMax(numbers)
	// fmt.Printf("data: %v\nMin: %d\nMax: %d\n", numbers, min, max)
	multireturn()
	spreadOperator()
	iife()       //immediately executed
	retClosure() //return function as Closure
	funcAsParam()
}

func retClosure() {

	findMax := func(numbers []int, max int) (int, func() []int) {
		var res []int

		for _, number := range numbers {
			if number <= max {
				res = append(res, number)
			}
		}

		return len(res), func() []int {
			return res
		}
	}

	var (
		max                 = 3
		numbers             = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
		howMany, getNumbers = findMax(numbers, max)
		theNumbers          = getNumbers()
	)

	fmt.Println("\nretClosure")
	fmt.Println("numbers\t", numbers)
	fmt.Printf("find \t: %d\n", max)
	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

func iife() {
	fmt.Println("\nIIFE")
	//IIFE karena ada kurung belakang
	var newNumbers = func(min int) []int {
		var filters []int
		for _, number := range numbers {
			if number < min {
				continue
			}
			filters = append(filters, number)
		}
		return filters
	}(3)

	fmt.Println("originalNumber: ", numbers)
	fmt.Println("filteredNumber: ", newNumbers)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func multireturn() {
	var calculateCircle = func(d float64) (float64, float64) {
		r := d / 2
		area := math.Pi * math.Pow(r, 2)
		circumference := math.Pi * d

		return area, circumference
	}

	var (
		d                   float64 = 15
		area, circumference float64 = calculateCircle(d)
	)

	fmt.Println("\nMultireturn")
	fmt.Printf("Luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("Keliling lingkaran\t: %.2f \n", circumference)
}

func spreadOperator() {
	var avg = func(numbers ...int) float64 {
		total := 0
		for _, number := range numbers {
			total += number
		}

		value := float64(total) / float64(len(numbers))
		return value
	}
	var myHobbies = func(name string, hobbies ...string) {
		var hobbiesAsString = strings.Join(hobbies, ", ")

		fmt.Printf("Hello my name is %s: \n", name)
		fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
	}

	average := avg(10, 12, -7, 12, 8, 9, 0, 4, -22)
	lycaasHobbies := []string{"walking", "playing", "sun-bathing"}
	fmt.Println("\nSpreadOperator")
	fmt.Printf("Avg: %.3f", average)
	myHobbies("lycaa", lycaasHobbies...)
	myHobbies("meyer", "reading", "sleeping", "barking", "playing with pops")
}

func funcAsParam() {
	var students = []string{"Lycaa", "Miko", "Miki", "Hime", "Mika", "Meyer", "Coy-coy"}

	var cekO = func(word string) bool {
		return strings.Contains(word, "o")
	}
	var cekLen = func(word string) bool {
		return len(word) == 5
	}
	var filter = func(words []string, callback func(string) bool) []string {
		var result []string

		for _, word := range words {
			if callback(word) {
				result = append(result, word)
			}
		}
		return result
	}

	fmt.Println("filter ada huruf \"o\"\t ", filter(students, cekO))
	fmt.Println("filter ada jumlah kata \"5\" ", filter(students, cekLen))

}
