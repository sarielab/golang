package main

import "fmt"

func main() {
	/*
	 ARRAY
	 - alokasi len dari awal or append
	 - kumpulan tipe data sama => nilai default
	 - deklarasi :
	    1. var hobbies [4]string
	    2. horizontal
	      var pergies = [2]string{"dyah", "aini", "lely", "wulan"}
	    3. vertikal
	      var hacktivs = [3]string{
	        "shab",
	        "iam",
	        "hakiem", //=> koma diakhir
	      }
	    4. tanpa jumlah elemen
	      var numbers = [...]int{2,3,4,5,6}
	    5. make
	      var fruits = make([]string, 2)
	      fruits[0] = "asd"
	    6. multidimensi
	      + subdimentsi boleh ga ada jumlah data

	*/
	var hobbies [4]string
	hobbies[0] = "coding"
	hobbies[1] = "shooting"
	hobbies[2] = "felting"
	hobbies[3] = "drawing"

	var pergies = [2]string{"dyah", "aini"}
	// fmt.Println("Jumlah pergies \t", len(pergies))
	// fmt.Println("Isi pergies \t", pergies)

	// hobbies[4] = "snorkeling"

	var numbers1 = [2][3]int{[3]int{3, 2, 1}, [3]int{4, 5, 6}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var numbers3 = [2][2][5]int{{{1, 2, 3, 4, 5}, {5, 10, 15, 20, 25}}, {{3, 6, 9, 12, 15}, {4, 8, 12, 16}}}
	fmt.Println("numbers3", numbers3)

	for i := 0; i < len(hobbies); i++ {
		fmt.Printf("hobby %d : %s\n", i, hobbies[i])
	}

	for i, pergi := range pergies {
		fmt.Printf("pergi %d : %s \n", i, pergi)
	}
	// harus pake i padahal kalau i ga dipake golang ngamuk
	// atau
	// for i, _:= range fruits {}
	// for i := range fruits {} => idx i doang
	for _, pergi := range pergies {
		fmt.Printf("pergi : %s \n", pergi)
	}
	fmt.Println("===============================")
	for i := range pergies {
		fmt.Println(i)
	}

}
