package main

import "fmt"

/*
  LOOP
    for => dengan args, kondisi saja, for
      break continue
      for range (array)
    foreach
    while
*/
func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("----------------------------")
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println("----------------------------")
	i = 3

	for {
		fmt.Println("Angka", i)

		i++
		if i == 5 {
			break
		}
	}

	fmt.Println("----------------------------")

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	//breakouterloop
changelable:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break changelable
			}
			fmt.Print("matriks [", i, "] [", j, "]", "\n")
		}
	}

}
