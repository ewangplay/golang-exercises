package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: base64tool <src_str>")
		os.Exit(1)
	}

	src_str := os.Args[1]

	fmt.Println("Origin String: ", src_str)

	base64_str := base64.StdEncoding.EncodeToString([]byte(src_str))

	fmt.Println("Base64 String: ", base64_str)

	restore_str, err := base64.StdEncoding.DecodeString(base64_str)
	if err != nil {
		fmt.Println("DecodeString fail.", err)
		os.Exit(1)
	}

	fmt.Println("Restore String: ", string(restore_str))
}
