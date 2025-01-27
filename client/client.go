package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {

	var url string
	var num string

	flag.StringVar(&url, "url", "http://localhost:8080/a", "url string flag")
	flag.StringVar(&num, "num", "12345678901234", "number flag in string")
	flag.Parse()

	c := http.Client{
		//Timeout: time.Second * 60,
	}
	req := getRequest(url, num)

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	x, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(x))

}

func getRequest(url string, num string) *http.Request {

	data := []byte(num)
	dataReader := bytes.NewReader(data)

	r, err := http.NewRequest(http.MethodPost, url, dataReader)

	if err != nil {
		fmt.Println("request creation error :", err.Error())
		return nil
	}
	return r
}
