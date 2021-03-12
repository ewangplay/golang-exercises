package main

import (
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"os"

	cl "github.com/ewangplay/cryptolib"
)

func main() {
	cfg := &cl.Config{
		ProviderName: "SW",
	}
	csp, err := cl.GetCSP(cfg)
	if err != nil {
		fmt.Printf("Get default CSP failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("ED25519 test ...")
	signAndVerify(csp, &cl.ED25519KeyGenOpts{})

	fmt.Println("ECDSA with default mode test ...")
	signAndVerify(csp, &cl.ECDSAKeyGenOpts{})

	fmt.Println("ECDSA with P224 mode test ...")
	signAndVerify(csp, &cl.ECDSAKeyGenOpts{Curve: elliptic.P224()})

	fmt.Println("ECDSA with P256 mode test ...")
	signAndVerify(csp, &cl.ECDSAKeyGenOpts{Curve: elliptic.P256()})

	fmt.Println("ECDSA with P384 mode test ...")
	signAndVerify(csp, &cl.ECDSAKeyGenOpts{Curve: elliptic.P384()})

	fmt.Println("ECDSA with P521 mode test ...")
	signAndVerify(csp, &cl.ECDSAKeyGenOpts{Curve: elliptic.P521()})
}

func signAndVerify(csp cl.CSP, opts cl.KeyGenOpts) error {
	k, err := csp.KeyGen(opts)
	if err != nil {
		fmt.Printf("KeyGen failed: %v\n", err)
		return err
	}

	kd, _ := k.Bytes()
	fmt.Println("Private Key: ", hex.EncodeToString(kd))

	digest := []byte("hello,world")
	signature, err := csp.Sign(k, digest)
	if err != nil {
		fmt.Printf("Sign failed: %v\n", err)
		return err
	}

	pubKey, err := k.PublicKey()
	if err != nil {
		fmt.Printf("Get public key failed: %v\n", err)
		return err
	}

	kd, _ = pubKey.Bytes()
	fmt.Println("Public Key: ", hex.EncodeToString(kd))
	fmt.Println("Signature: ", hex.EncodeToString(signature))

	valid, err := csp.Verify(pubKey, digest, signature)
	if err != nil {
		fmt.Printf("Verify failed: %v\n", err)
		return err
	}
	if !valid {
		fmt.Println("Signature verification failed.")
		return fmt.Errorf("signature verification failed")
	}

	fmt.Println("The signature is validated successfully.")
	return nil
}
