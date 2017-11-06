package main

/*
  Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang
  dibungkus dengan nama tertentu.
  Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map , hanya saja key-nya
  sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.
  Dengan memanfaatkan struct, data akan terbungkus lebih rapi dan mudah di-maintain.
  Struct merupakan cetakan, digunakan untuk mencetak variabel objek (istilah untuk variabel
  yang memiliki property). Variabel objek memiliki behaviour atau sifat yang sama sesuai
  struct pencetaknya. Konsep ini sama dengan konsep class pada pemrograman berbasis
  objek. Sebuah buah struct bisa dimanfaatkan untuk mencetak banyak objek.

  - Deklarasi
    var s1 = student {}
    s1.name = "a"
    var s1 = student{"joana", 4, []string{"reading webtoon", "watching disney channel"}}
    var s1 student
    s1.name ="b"
  - WHY pointer in struct
    bisa lsg tunjuk tanpa di-dereferensikan
  - Embedded Struct
    simpan hasil struct ke property struct lain
    + mutable
    + bisa lsg akses ke anak tanpa .parent pas get/set
    + bisa sama propsname student.age VS student.person.age pas di get tetep beda
  - Anonymous Struct
    + utk struct yg dipake 1x aja
    + ditampung di var lsg

*/

import (
	"fmt"
)

type student struct {
	name    string   `tag1`
	grade   int      `tag2`
	hobbies []string `tag3`
}

type pet struct {
	owner string
	dog
}

type dog struct {
	name       string
	race       string
	hair_color string
}

type stud struct {
	profile struct {
		name  string
		level int
	}
	grade   int
	hobbies []string
}

func main() {
	struct1()
	structPointer()
	embeddedStruct()
	anonStruct()
	sliceStruct()
}

func struct1() {
	var s1 student
	s1.name = "Lycaa"
	s1.grade = 4
	s1.hobbies = []string{"Walking", "Dancing", "Barking"}

	var s2 = student{"Meyer", 2, []string{"Biting"}}
	var s3 = student{
		name:    "Mii",
		grade:   3,
		hobbies: []string{"Playing", "Walking", "Bathing"}}
	fmt.Println(s3)
	fmt.Println("name: ", s1.name)
	fmt.Println("grade: ", s1.grade)
	fmt.Println("hobbies: ", s1.hobbies)
	fmt.Println(s2)
}

func structPointer() {
	var s1 = student{"Meyer", 2, []string{"Biting"}}
	var s2 *student = &s1

	fmt.Println("\nStruct Pointer")
	fmt.Println("student1, name", s1.name)
	fmt.Println("student2, name", s2.name)
	s1.name = "Meyer berubah nama"
	fmt.Println("student1, name", s1.name)
	fmt.Println("student2, name", s2.name)
	s2.name = "Ethan"
	fmt.Println("student1, name", s1.name)
	fmt.Println("student2, name", s2.name)
}

func embeddedStruct() {
	var d1 = pet{owner: "Pops"}
	d1.name = "Lycaa"
	d1.race = "Toy poodle"
	d1.hair_color = "white"

	fmt.Println("\nEmbeddedStruct")
	fmt.Println("owner: ", d1.owner)
	fmt.Println("name", d1.name)
	fmt.Println("name", d1.dog.name)
}

func anonStruct() {
	var pets = struct {
		owner string
		dog
	}{}

	var pets2 = struct {
		owner string
		dog
	}{
		owner: "Wen",
		dog:   dog{"Miki", "Bastard", "Black"},
	}

	pets.owner = "Poppy"
	pets.dog = dog{"Meyer", "Kintamani", "white"}
	fmt.Println("\nAnonStruct")
	fmt.Println(pets)
	fmt.Println(pets2)
}

func sliceStruct() {
	var allDogs = []dog{
		{name: "Lycaa", race: "Toy Poodle", hair_color: "white"},
		{name: "Miko", race: "Bastard", hair_color: "gray"},
	}
	allDogs = append(allDogs, dog{"Meyer", "Kintamani", "white"})

	fmt.Println(allDogs)
}
