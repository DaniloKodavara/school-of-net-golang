package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	name = "TESTE"
)

func main() {

	res, _ := http.Get("http://google.com")
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s", body)

}
