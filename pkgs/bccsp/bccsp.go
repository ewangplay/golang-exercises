package main

import (
	//"bytes"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
)

func main() {
	// New SW BCCSP instance
	config := &factory.FactoryOpts{
		ProviderName: "SW",
		SwOpts: &factory.SwOpts{
			HashFamily: "SHA2",
			SecLevel:   256,
			Ephemeral:  false,
			FileKeystore: &factory.FileKeystoreOpts{
				KeyStorePath: "./keystore",
			},
		},
	}
	csp, err := factory.GetBCCSPFromOpts(config)
	if err != nil {
		fmt.Printf("Failed to get sw bccap: %v\n", err)
		return
	}

	// Generate ecdsa private key
	priKey, err := csp.KeyGen(&bccsp.ECDSAKeyGenOpts{Temporary: false})
	if err != nil {
		fmt.Printf("Failed to generate ecdsa private key: %v\n", err)
		return
	}
	fmt.Printf("Symmetric: %v\n", priKey.Symmetric())
	fmt.Printf("Private: %v\n", priKey.Private())

	// Retrieve ecdsa public key
	pubKey, err := priKey.PublicKey()
	if err != nil {
		fmt.Printf("Failed to get public key: %v\n", err)
	}
	fmt.Printf("Symmetric: %v\n", pubKey.Symmetric())
	fmt.Printf("Private: %v\n", pubKey.Private())

	// Sign message
	msg := []byte("hello, world.")
	msgSign, err := csp.Sign(priKey, msg, nil)
	if err != nil {
		fmt.Printf("Failed sign message: %v\n", err)
		return
	}

	// Verify the signature
	valid, err := csp.Verify(pubKey, msgSign, msg, nil)
	if err != nil {
		fmt.Printf("Failed verify signature: %v\n", err)
		return
	}
	fmt.Println(func(valid bool) string {
		if valid {
			return "Verify signature passed!"
		} else {
			return "Verify signautre not passed!"
		}
	}(valid))

	// 加解密的测试没有通过，还需要继续研究
	/*
		// Encrypt message
		cipherMsg, err := csp.Encrypt(pubKey, msg, nil)
		if err != nil {
			fmt.Printf("Failed encrypt message: %v\n", err)
			return
		}

		// Decrypt the cipher message
		plainMsg, err := csp.Decrypt(priKey, cipherMsg, nil)
		if err != nil {
			fmt.Printf("Failed decrypt cipher message: %v\n", err)
			return
		}
		if bytes.Compare(plainMsg, msg) == 0 {
			fmt.Println("decryption succ.")
		} else {
			fmt.Println("decryption fail.")
		}
	*/
}
