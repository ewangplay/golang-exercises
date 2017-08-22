package main

import (
	"fmt"
	"strings"
)

func main() {
	result := make(map[string]interface{})

	result["code"] = 0
	result["id"] = 123
	result["hash"] = "sjdfjdsfjsdfkdsjfksdjfdjs"
	result["data"] = map[string]interface{}{
		"name":     "Tom",
		"type":     1,
		"order_id": "456",
	}

	ret := formatResult(result)

	fmt.Printf("%s\n", ret)
}

func formatResult(result map[string]interface{}) string {
	ret := "{"

	for k, v := range result {
		switch v.(type) {
		case string:
			ret += fmt.Sprintf("\"%s\":\"%s\",", k, v)
		case map[string]interface{}:
			ret += fmt.Sprintf("\"%s\":%s,", k, formatResult(v.(map[string]interface{})))
		default:
			ret += fmt.Sprintf("\"%s\":%v,", k, v)
		}
	}
	if strings.HasSuffix(ret, ",") {
		ret = strings.TrimRight(ret, ",")
	}
	ret += "}"

	return ret
}
