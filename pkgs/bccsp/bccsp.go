package main

import (
	"bytes"
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

	// Sign / Verify test using ECDSA
	fmt.Println("Sign / Verify test using ECDSA...")
	data := []byte("hello, world.")
	err = signVerify(csp, data)
	if err != nil {
		fmt.Printf("Sign / Verify testing failed: %v\n", err)
		return
	}

	// Encrypt / Encrypt test using AES
	fmt.Println("Encrypt / Encrypt test using AES...")
	err = encryptDecrypt(csp, data)
	if err != nil {
		fmt.Printf("Encrypt / Decrypt testing failed: %v\n", err)
		return
	}

}

func signVerify(csp bccsp.BCCSP, data []byte) error {
	var err error

	// Generate ecdsa private key
	priKey, err := csp.KeyGen(&bccsp.ECDSAKeyGenOpts{Temporary: false})
	if err != nil {
		fmt.Printf("Failed to generate ecdsa private key: %v\n", err)
		return err
	}
	fmt.Printf("Symmetric: %v\n", priKey.Symmetric())
	fmt.Printf("Private: %v\n", priKey.Private())

	// Retrieve ecdsa public key
	pubKey, err := priKey.PublicKey()
	if err != nil {
		fmt.Printf("Failed to get public key: %v\n", err)
		return err
	}
	fmt.Printf("Symmetric: %v\n", pubKey.Symmetric())
	fmt.Printf("Private: %v\n", pubKey.Private())

	// Sign message
	msgSign, err := csp.Sign(priKey, data, nil)
	if err != nil {
		fmt.Printf("Failed sign message: %v\n", err)
		return err
	}

	// Verify the signature
	valid, err := csp.Verify(pubKey, msgSign, data, nil)
	if err != nil {
		fmt.Printf("Failed verify signature: %v\n", err)
		return err
	}
	fmt.Println(func(valid bool) string {
		if valid {
			return "Verify signature passed!"
		}
		return "Verify signautre not passed!"
	}(valid))

	return nil
}

func encryptDecrypt(csp bccsp.BCCSP, data []byte) error {
	// Generate aes key
	key, err := csp.KeyGen(&bccsp.AESKeyGenOpts{Temporary: false})
	if err != nil {
		fmt.Printf("Failed to generate ecdsa private key: %v\n", err)
		return err
	}
	fmt.Printf("Symmetric: %v\n", key.Symmetric())
	fmt.Printf("Private: %v\n", key.Private())

	// Encrypt message
	cipherMsg, err := csp.Encrypt(key, data, &bccsp.AESCBCPKCS7ModeOpts{})
	if err != nil {
		fmt.Printf("Failed encrypt message: %v\n", err)
		return err
	}

	// Decrypt the cipher message
	plainMsg, err := csp.Decrypt(key, cipherMsg, &bccsp.AESCBCPKCS7ModeOpts{})
	if err != nil {
		fmt.Printf("Failed decrypt cipher message: %v\n", err)
		return err
	}

	if bytes.Compare(plainMsg, data) == 0 {
		fmt.Println("decryption succ.")
	} else {
		fmt.Println("decryption fail.")
	}

	return nil
}
