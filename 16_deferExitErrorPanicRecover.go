package main

/*
  defer => stlh semua non defer dijalankan kec return LIFO (aka dari bawah keatas)
  exit(error int) => STOP bener2 bahkan defer
  error => tipedata, prop method : Error
    bbrp balikin error (strconv)
  panic => fungsi setelahnya gjln kecuali defer
  recover => handle panic error
    better taro di defer
*/

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func defer1() {
	fmt.Println("\n\n==============Defer==============")
	var orderSomeFood = func(menu string) {
		defer fmt.Println("Terimakasih, silahkan tunggu")

		if menu == "indomie" {
			fmt.Println("Pilihan tepat! ")
			fmt.Println("Indomie kami yang terenak men")
			return
		}

		fmt.Println("Pesanan anda: ", menu)
	}

	orderSomeFood("cake")
	orderSomeFood("pizza")
}

func exit1() {
	fmt.Println("\n\n==============Exit==============")
	var orderSomeFood = func(menu string) {
		defer fmt.Println("Terimakasih, silahkan tunggu")
		os.Exit(1)
		if menu == "indomie" {
			fmt.Println("Pilihan tepat! ")
			fmt.Println("Indomie kami yang terenak men")
			return
		}

		fmt.Println("Pesanan anda: ", menu)
	}

	orderSomeFood("cake")
	orderSomeFood("pizza")
}

func error1() {
	fmt.Println("\n\n=============Error=============")
	num_tes := "1"
	str_tes := "tes"

	var num1, num2 int
	var err1, err2 error

	num2, err2 = strconv.Atoi(str_tes)
	num1, err1 = strconv.Atoi(num_tes)

	if err2 == nil {
		fmt.Println(num2, "is number")
	} else if err2 != nil {
		fmt.Println(num2, "is not a number")
		fmt.Println(err2.Error())
	}

	if err1 == nil {
		fmt.Println(num1, "is number")
	} else if err1 != nil {
		fmt.Println(num1, "is not a number")
		fmt.Println(err1.Error())
	}
}

func error2() {
	fmt.Println("\n\n=============Error Custom=============")

	var validate = func(input string) (bool, error) {
		if strings.TrimSpace(input) == "" {
			return false, errors.New("can not be empty")
		}
		return true, nil
	}

	var _, err = validate("                              ")

	fmt.Println(err)
}

var validate = func(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("can not be empty")
	}
	return true, nil
}

func panic1() {
	fmt.Println("\n\n=============Panic=============")

	defer fmt.Println("cobain defer setelah panik")

	if valid, err := validate("                              "); valid {
		fmt.Println("Halo, ", valid)
	} else {
		panic(err.Error())
	}
}

func recover1() {
	fmt.Println("\n\n=============Recover=============")

	var catch = func() {
		if r := recover(); r != nil {
			fmt.Println("Error occured", r)
		} else {
			fmt.Println("Apps running perfectly")
		}
	}

	defer catch()

	if valid, err := validate("                              "); valid {
		fmt.Println("Halo, ", valid)
	} else {
		panic(err.Error())
	}
}

func main() {
	defer1()
	error1()
	error2()
	recover1()
	panic1()
	exit1()
}
