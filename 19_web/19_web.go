package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar")
}

func main() {
	urlParsing()
	jsonParsing()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "Poppy Sari",
			"Message": "Have a nice day",
		}
		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func urlParsing() {
	var urlString = "http://tes.com:80/search?genre=horror&completed=false"
	var u, err = url.Parse(urlString)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)
	fmt.Println("protocol: ", u.Scheme) //http
	fmt.Println("host: ", u.Host)       // tes.com:80
	fmt.Println("path: ", u.Path)       //search

	for _, query := range u.Query() {
		fmt.Println(query)
	}
}

func jsonParsing() {
	type User struct {
		FullName string `json:"Name"`
		Age      int
	}

	var jsonString = `[{"Name":"John Wick", "Age": 27}, {"Name":"Lilac", "Age": 21}]`
	//var jsonString = `{"Name":"John Wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	fmt.Println("\n\n========JSON PARSING========")

	var toStruct = func() {
		var data []User
		//var data User
		var err = json.Unmarshal(jsonData, &data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//data.FullName
		fmt.Println("user : ", data[0].FullName)
		fmt.Println("age: ", data[0].Age)
	}

	var toInterface = func() {
		var data1 []map[string]interface{}
		//var data1 map[string]interface{}

		json.Unmarshal(jsonData, &data1)
		fmt.Println(data1)
		fmt.Println("user: ", data1[0]["Name"])
		fmt.Println("age: ", data1[0]["Age"])
		// fmt.Println("age: ", data1["Age"])
	}

	var toArrayObject = func() {
		var data []User
		var err = json.Unmarshal([]byte(jsonString), &data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("user : ", data[1].FullName)

	}

	var encodeJSON = func() {
		var object = []User{{"Lina", 22}, {"Lycaa", 17}}
		var jsonData, err = json.Marshal(object)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var jsonString = string(jsonData)
		fmt.Println(jsonString)
	}

	toStruct()
	toInterface()
	toArrayObject()
	encodeJSON()
}
