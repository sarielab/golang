package main

import (
	"fmt"
	"os"
	"time"
)

//page 172
/*
  time.Sleep => suru tunggu
  time.NewTimer() => balikin object *time.Timer
		var	timer	=	time.NewTimer(4	*	time.Second)
		fmt.Println("start")
		<-timer.C
		fmt.Println("finish")
  time.After vs time.Sleep => balikin data channel
    <- time.After(...)
*/
func time1() {
	var time1 = time.Now()
	fmt.Printf("time1 %v \n", time1) // time1 2017-10-10 17:57:44.092993804 +0700 WIB m=+0.000183982
	fmt.Printf("month %v year %v weekday (%v %v) \n", time1.Month(), time1.Year(), time1.Weekday(), time1.Weekday().String())
	fmt.Println("tahun minggu ke mulai awal taon")
	fmt.Println(time1.ISOWeek())

	var time2 = time.Date(2017, 07, 11, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v \n", time2)
	// time2 2017-07-11 10:20:00 +0000 UTC
}

func parsingTime() {
	fmt.Println("============parsingTime==============")

	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02	15:04:05"
	value = "2015-09-02	08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())
	//	2015-09-02	08:04:00	+0000	UTC

	layoutFormat = "02/01/2006	MST"
	value = "02/09/2015	WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())
}

func passingPredefinedTime() {
	var date, _ = time.Parse(time.RFC822, "02	Sep	15	08:00	WIB")
	var dateS1 = date.Format("Monday	02,	January	2006	15:04	MST")
	fmt.Println("dateS1", dateS1)
	//	Wednesday	02,	September	2015	08:00	WIB
	var dateS2 = date.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)
	//	2015-09-02T08:00:00+07:00
}

func timer1() {
	fmt.Println("start")
	fmt.Println(time.Second)
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 second")
}

func newTimer() {
	/*
	  Statement	 	var	timer	=	time.NewTimer(4	*	time.Second)	 	mengindikasikan	bahwa	nantinya
	  akan	ada	data	yang	dikirimkan	ke	channel	 	timer.C	 	setelah	4	detik	berlalu.	Baris	kode	 	<-
	  timer.C	 	menandakan	penerimaan	data	dari	channel	 	timer.C	 .	Karena	penerimaan	channel
	  sendiri	sifatnya	adalah	blocking,	maka	statement	 	fmt.Println("finish")	 	baru	akan
	  dieksekusi	setelah	4	detik.
	*/
	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")
}

func afterFunc() {
	/*
	  NOTE:
	  IF NO channel serahTerima =>	eksekusi time.AfterFunc()	=> asynchronous	& non-blocking
	  ELSE	fungsi	async	& non	blocking	sampe baris penerimaan code
	*/
	fmt.Println("\n\n===================After Func==============")
	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")

}

func after() {
	fmt.Println("\n\n===================After==============")
	<-time.After(4 * time.Second)
	fmt.Println("expired")
}

func main() {
	// time1()
	// parsingTime()
	// passingPredefinedTime()
	//
	// //=========startTimer
	// timer1()
	// newTimer()
	// afterFunc()
	// after()

	//===========function X goroutine
	var timer = func(timeout int, ch chan<- bool) {
		time.AfterFunc(time.Duration(timeout)*time.Second, func() {
			ch <- true
		})
	}

	var watcher = func(timeout int, ch <-chan bool) {
		<-ch
		fmt.Println("yah waktunya abis deh :((")
		os.Exit(0)
	}

	var ch = make(chan bool)
	var input string

	fmt.Println("1 + 1 adalah...")
	fmt.Scan(&input)

	go timer(4, ch)
	go watcher(4, ch)

	if input == "2" {
		fmt.Println("pintar juga anda")
	} else {
		fmt.Println("tk nya sogok ya? :P")
	}

	fmt.Scan(&input)
}
