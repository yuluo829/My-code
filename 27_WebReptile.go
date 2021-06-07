package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	url := "https://peapix.com/"
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error",err)
		return
	}
	fmt.Println(string(body))
}
