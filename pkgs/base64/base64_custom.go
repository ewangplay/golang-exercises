package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

const (
	/*
	 *通过这个字符集生成的Base64 Encoding对象相当于base64.StdEncoding对象
	 *跟直接使用base64.StdEncoding的效果一样
	 */
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

	/*
	 *通过这个字符集生成的Base64 Encoding对象不能成功还原编码
	 */
	//base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: base64tool <src_str>")
		os.Exit(1)
	}

	src_str := os.Args[1]

	fmt.Println("Origin String: ", src_str)

	base64Coder := base64.NewEncoding(base64Table)

	base64_str := base64Coder.EncodeToString([]byte(src_str))

	fmt.Println("Base64 String: ", base64_str)

	restore_str, err := base64Coder.DecodeString(base64_str)
	if err != nil {
		fmt.Println("DecodeString fail.", err)
		os.Exit(1)
	}

	fmt.Println("Restore String: ", string(restore_str))
}
