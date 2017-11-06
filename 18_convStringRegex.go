package main

import (
	"fmt"
	"strconv"
)

/*
  strconv
    - Atoi int to str
    - itoa str to int

*/
// strToNum2    = "asdasd"
func conv() {
	var (
		strToNum     = "123"
		numToStr     = 123
		num, err     = strconv.Atoi(strToNum)
		str          = strconv.Itoa(numToStr)
		formatint    = int64(numToStr)
		strFormatInt = strconv.FormatInt(formatint, 8)
		num2, err2   = strconv.ParseInt(str, 10, 64)
	)

	if err == nil {
		fmt.Println(num)
	}

	if err2 == nil {
		fmt.Println(num2)
	}

	fmt.Println(str)
	fmt.Println(strFormatInt)
}

func stringFun() {

}

func regexing() {

}

func main() {
	conv()
	stringFun()
	regexing()
}
