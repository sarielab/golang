package main

/*"
  Interface	kosong	atau	 	interface{}	 	adalah	tipe	data	yang	sangat	spesial.	Variabel	bertipe
  ini	bisa	menampung	segala	jenis	data,	bahkan	array,	bisa	pointer	bisa	tidak	(konsep	ini
  disebut	dengan	dynamic	typing).
*/

import (
	"fmt"
	"strings"
)

var secret interface{}
var data map[string]interface{}

func main() {
	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, " multiplied by 10 is is ", number)

	secret = "Ethan Hunt"
	fmt.Println(secret)

	secret = []string{"apple", "pear", "mango"}
	var breakfasts = strings.Join(secret.([]string), ", ")
	fmt.Println(breakfasts)

	secret = 12.4
	fmt.Println(secret)

	data = map[string]interface{}{
		"name":      "Ethan",
		"grade":     2,
		"breakfast": []string{"egg", "coffee"},
	}

	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23, "hobbies": []string{"tidur", "mati"}},
		{"name": "Ethan", "age": 23, "hobbies": []string{"main", "mandi"}},
		{"name": "Bourne", "age": 22, "hobbies": []string{"catur", "main"}},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age	is", each["age"], "hobbies are ", each["hobbies"])
	}

}
