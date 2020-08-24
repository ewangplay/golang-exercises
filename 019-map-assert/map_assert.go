package main

import (
	"fmt"
)

func main() {
	result := make(map[interface{}]interface{})

	result["list"] = make([]map[string]string, 10)

	v, ok := result["list"].([]map[string]string)
	if ok {
		for i := 0; i < 10; i++ {
			v[i] = make(map[string]string)

			key := fmt.Sprintf("key%v", i+1)
			value := fmt.Sprintf("value%v", i+1)

			v[i][key] = value
		}
	}

	fmt.Println(result)
}
