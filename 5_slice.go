package main

import "fmt"

var pets = []string{"iguana", "bunny", "dog", "cat", "wombat"}

func main() {
	/*
	 ARRAY VS SLICE
	 - slice ga didefinisiin indexnya
	 - slice reference -> ubah idx lama = ubah baru

	 Slice merupakan tipe data reference atau referensi.
	 Artinya jika ada slice baru yang terbentuk dari slice lama,
	 maka data elemen slice yang baru akan memiliki alamat memori
	 yang sama dengan elemen slice lama.
	 Setiap perubahan yang terjadi di elemen slice baru,
	 akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.


	 SLICE
	  - len() digunakan untuk menghitung jumlah elemen slice yang ada
	  - cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice
	  - append() digunakan untuk menambahkan elemen pada slice.
	      len == cap => elemen hasil append = ref baru
	      len < cap => dia liat dari ref asli trus timpa index ke segitu
	  - copy() digunakan untuk men-copy elemen slice pada parameter ke-2,
	      untuk digabungkan dengan slice pada parameter ke-1
	  - 3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya.
	      Cara menggunakannnya yaitu dengan menyisipkan angka kapasitas di belakang,
	      seperti fruits[0:1:1].
	      Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice
	      yang akan di slicing.

	*/
	// slicing()
	// appending()
	// capAppending()
	// copying()
	slicing3Idx()
}

func slicing() {
	fmt.Println("===================SLICING================")
	fmt.Printf("pets[:] \t%s\n", pets[:])
	fmt.Printf("pets[0:2] \t%s\n", pets[0:2])
	fmt.Printf("pets[2:] \t%s\n", pets[2:])
	fmt.Printf("pets[:2] \t%s\n", pets[:2])
	fmt.Printf("pets[2:] \t%s\n", pets[2:])

	newPets := pets[0:2]
	newPets[1] = "alpacca"

	fmt.Printf("slice is %s %d\n", newPets, len(newPets))
	fmt.Printf("cap %d \n", cap(newPets))
}

func appending() {
	fmt.Println("===================APPENDING===============")
	var appendPet = append(pets, "bat", "capybara")
	fmt.Printf("append pets \t: %s\n", appendPet)
	fmt.Printf("pets \t\t: %s\n", pets)
}

func capAppending() {
	fmt.Println("===============CAPAPPENDING=================")
	bPets := pets[0:2]

	fmt.Println(cap(bPets)) //5
	fmt.Println(len(bPets)) //2

	fmt.Println(pets)
	fmt.Println(bPets)
	fmt.Println("-----append--------")
	var cPets = append(bPets, "lycaa")
	fmt.Println(pets)  // timpa dog -> lycaa
	fmt.Println(bPets) // ga ada perubahan
	fmt.Println(cPets) // dah ter append
}

func copying() {
	var hobbies = []string{"coding", "reading", "observing"}
	var copiedElement = copy(pets, hobbies) //hobbies timpa pet

	fmt.Println("==================COPYING====================")
	fmt.Println(pets)
	fmt.Println(hobbies)
	fmt.Println(copiedElement)
}

func slicing3Idx() {
	bPets := pets[0:2]
	cPets := pets[0:2:2]

	fmt.Println("==================== 3 IDX ====================")

	fmt.Println(pets)
	fmt.Println(len(pets))
	fmt.Println(cap(pets))

	fmt.Println("pets[0:2]")
	fmt.Println(bPets)
	fmt.Println(len(bPets))
	fmt.Println(cap(bPets))

	fmt.Println("pets[0:2:2]")
	fmt.Println(cPets)
	fmt.Println(len(cPets))
	fmt.Println(cap(cPets))
}
