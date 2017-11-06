package main

import (
	"fmt"
	"runtime"
)

/*
  channel utk hubungkan go routine2
  1. var newChannel = make(chan tipedata)
  2. assign blocking sementara
    newChannel <- data
    var gr = <- newChannel
  3. tipe
    - unbuffered aka DEFAULT
      Ketika
       - ada	goroutine kirim	data	via	channel,
       - ada	goroutine	lain	yang	bertugas	menerima
      ==> data	dari	channel	yang	sama,	dengan	proses	serah-terima	yang	bersifat	blocking.
      ==> tunggu fin
    - buffered pg 145
     - pas make => msgs := make(chan int, 2) => 0 1 2 (3 proses)
  4. channelDirection
    - ch chan string    => kirim terima
    - ch <-chan string  => terima aja
    - ch chan<-string   => kirim aja
  5. for - range - closure
    for i := 0; i < 20 ; i++ { ch <- fmt.Sprintf("data %d", i)}
    for msg := range ch {fmt.Println(msg)}
    close(ch)

*/

func main() {
	/*
		    UNBUFFERED CHANNEL AKA DEFAULT

		 		  // runtime.GOMAXPROCS(2)
					// var msgs = make(chan string)
					// var getMsg = func(ppl string) {
					// 	msgs <- "hello " + ppl
					// }
		      //
					// go getMsg("a")
					// go getMsg("b")
					// go getMsg("c")
		      //
					// var msg1 = <-msgs
					// fmt.Println(msg1)
					// var msg2 = <-msgs
					// fmt.Println(msg2)
					// var msg3 = <-msgs
					// fmt.Println(msg3)


			// runtime.GOMAXPROCS(2)
		  //
			// var messages = make(chan string)
			// var printMessage = func(msg chan<-string) {
			// 	fmt.Println(<-msg)
			// }
		  //
			// for _, each := range []string{"Lycaa", "Meyer", "Miko"} {
			// 	go func(who string) {
			// 		var data = fmt.Sprintf("hello %s", who)
			// 		messages <- data
			// 	}(each)
			// }
		  //
			// for i := 0; i < 3; i++ {
			// 	printMessage(messages)
			// }

	*/
	/*
	   //BUFFERED

	   runtime.GOMAXPROCS(2)
	   messages := make(chan int, 2)
	   go func() {
	   for {
	     i := <-messages
	     fmt.Println("receive	data", i)
	   }
	   }()

	   for i := 0; i < 5; i++ {
	     fmt.Println("send	data", i)
	     messages <- i
	   }
	*/
	runtime.GOMAXPROCS(2)

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 11, 13, 17, 3}
	fmt.Println("numbers", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan float64)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %.2f \n", max)
		}
	}
}

func getAverage(numbers []int, ch chan<- float64) {
	var tot = 0
	var len = len(numbers)

	for _, number := range numbers {
		tot += number
	}

	ch <- float64(tot) / float64(len)
}

func getMax(numbers []int, ch chan<- float64) {
	var max = numbers[0]

	for _, number := range numbers {
		if max < number {
			max = number
		}
	}

	ch <- float64(max)
}
