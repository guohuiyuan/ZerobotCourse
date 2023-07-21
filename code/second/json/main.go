package main

import (
	"encoding/json"
	"fmt"

	"github.com/FloatTech/floatbox/web"
)

var (
	requestURL = "https://api.shadiao.app/chp"
)

type shadiao struct {
	Data struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"data"`
}

func main() {
	data, err := web.RequestDataWith(web.NewDefaultClient(), requestURL, "GET", "", web.RandUA(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	var s shadiao
	json.Unmarshal(data, &s)
	fmt.Printf("%+v\n", s)
}
