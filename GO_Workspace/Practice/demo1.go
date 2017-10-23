// demo1
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://api.openweathermap.org/data/2.5/weather?q=London&APPID=7e9d8941957867aa941c4c7873a3d8c8"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("postman-token", "59ce542a-ab2c-f481-b00d-6cf2cb3cf831")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
