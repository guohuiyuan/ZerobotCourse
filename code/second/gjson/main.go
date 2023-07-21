package main

import (
	"fmt"

	"github.com/FloatTech/floatbox/web"
	"github.com/tidwall/gjson"
)

var (
	requestURL = "https://api.shadiao.app/chp"
)

func main() {
	data, err := web.RequestDataWith(web.NewDefaultClient(), requestURL, "GET", "", web.RandUA(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	fmt.Println(gjson.GetBytes(data, "data.text").String())
}
