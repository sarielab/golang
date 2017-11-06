/*
  Race	condition	=>kondisi	dimana	lebih	dari	satu	goroutine,	mengakses	data	yang	sama
  pada	waktu	yang	bersamaan	(benar-benar	bersamaan).	Ketika	hal	ini	terjadi,	nilai	data
  tersebut	akan	menjadi	kacau.	Dalam	concurrency	programming	situasi	seperti	ini	ini
  sering	terjadi.

  Mutex	melakukan	pengubahan	level	akses	sebuah	data	menjadi	eksklusif,	menjadikan	data
  tersebut	hanya	dapat	dikonsumsi	(read	/	write)	oleh	satu	buah	goroutine	saja.	Ketika	terjadi
  race	condition,	maka	hanya	goroutine	yang	beruntung	saja	yang	bisa	mengakses	data
  tersebut.	Goroutine	lain	(yang	waktu	running	nya	kebetulan	bersamaan)	akan	dipaksa	untuk
  menunggu,	hingga	goroutine	yang	sedang	memanfaatkan	data	tersebut	selesai.
  Golang	menyediakan	 	sync.Mutex	 	yang	bisa	dimanfaatkan	untuk	keperluan	lock	dan
  unlock	data.
*/

package main

func main() {}
