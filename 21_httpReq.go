package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type studentDt struct {
	ID    string
	Name  string
	Grade int
}

const baseURL = "http://localhost:8080"

func fetchUsers() ([]studentDt, error) {
	var err error
	var client = &http.Client{}
	var data []studentDt

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUser(ID string) (studentDt, error) {
	var err error
	var client = &http.Client{}
	var data studentDt

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	req, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		return data, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return data, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}
func httpReq() {
	var users, err = fetchUsers()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, stu := range users {
		fmt.Println("data stu: ", stu)
	}
}

func httpReq2() {
	var user1, err = fetchUser("E001")
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	fmt.Println(user1)
}

func main() {
	httpReq()
	httpReq2()
}
