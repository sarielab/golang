package main

/*
  Goroutine	mirip	dengan	thread	thread,	tapi	sebenarnya	bukan.	Sebuah	native	thread	bisa
  berisikan	sangat	banyak	goroutine.	Mungkin	lebih	pas	kalau	goroutine	disebut	sebagai	mini
  thread.	Goroutine	sangat	ringan,	hanya	dibutuhkan	sekitar	2kB	memori	saja	untuk	satu
  buah	goroutine.	Ekseksui	goroutine	bersifat	asynchronous,	menjadikannya	tidak	saling
  tunggu	dengan	goroutine	lain.

  Untuk	menerapkan	goroutine,	proses	yang	akan	dieksekusi	sebagai	goroutine	harus
  dibungkus	kedalam	sebuah	fungsi.	Pada	saat	pemanggilan	fungsi	tersebut,	ditambahkan
  keyword	 	go	 	didepannya,	dengan	itu	goroutine	baru	akan	dibuat	dengan	tugas	adalah
  menjalankan	proses	yang	ada	dalam	fungsi	tersebut.


  HOW?
    1. runtime.GOMAXPROCS(jml)
    2. go namaFunction()
*/

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var print = func(till int, msg string) {
		for i := 0; i < till; i++ {
			fmt.Println((i + 1), msg)
		}
	}

	go print(5, "halo")
	print(5, "Apa kabar")

	var input string
	fmt.Scanln(&input) //ga slese sampe enter

}
