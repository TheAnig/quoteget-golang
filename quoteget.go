package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Quote, Author, Category string
}


func GetRequest(url string) *http.Response {
	res, err := http.Get("https://talaikis.com/api/quotes/random")

	if err != nil {
		log.Fatal(err)
	}

	return res
}

func ReadBody(res *http.Response) string {
	byteArr, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(byteArr)
}

func ReadJSON(s string) Message {
	var m Message

	dec := json.NewDecoder(strings.NewReader(s))

	err := dec.Decode(&m)

	if err != nil {
		log.Fatal(err)
	}

	return m
}



func main() {

	res := GetRequest("https://talaikis.com/api/quotes/random")

	s := ReadBody(res)

	m := ReadJSON(s)

	fmt.Printf("\n Quote : %s \n\n Author : %s \n\n", m.Quote, m.Author)

}