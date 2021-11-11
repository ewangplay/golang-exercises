package main

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/pkcs12"
)

func main() {
	pfxData, err := ioutil.ReadFile("./test.pfx")
	if err != nil {
		fmt.Println("ReadFile failed: ", err)
		return
	}

	pemBlocks, err := pkcs12.ToPEM(pfxData, "111111")
	if err != nil {
		fmt.Println("Decode pfx data failed: ", err)
		return
	}

	var rsaPrivateKey *rsa.PrivateKey
	var certs []*x509.Certificate
	for _, b := range pemBlocks {
		privateKey, err := x509.ParsePKCS8PrivateKey(b.Bytes)
		if err == nil && privateKey != nil {
			fmt.Println("ParsePKCS8PrivateKey Succ!")
			rsaPrivateKey = privateKey.(*rsa.PrivateKey)
			continue
		}
		certs, err = x509.ParseCertificates(b.Bytes)
		if err == nil {
			fmt.Println("ParseCertificates Succ!")
			continue
		}
	}
	if rsaPrivateKey == nil || certs == nil {
		fmt.Println("Parse key failed!")
		return
	}

}
