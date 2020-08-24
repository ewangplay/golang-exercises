package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := "/v(\\d)/(contents|mail|sms)/"
	url := "http://api.jiuzhilan.com/v1/contents/?act=add&type=type1&addr=beijing&token=12345678"
	regexp, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("comple regular expression fail. %v", err)
		return
	}

	matchs := regexp.FindStringSubmatch(url)
	if matchs == nil {
		fmt.Printf("not any match.\n")
		return
	}

	for _, str := range matchs {
		fmt.Printf("%v\n", str)
	}
}
