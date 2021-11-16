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
	signAndVerify(csp, &cl.ED25519KeyGenOpts{}, nil)

	fmt.Println("ECDSA with default mode test ...")
	signAndVerify(csp, &cl.ECDSAKeyGenOpts{}, nil)

	fmt.Println("ECDSA with P224 mode test ...")
	signAndVerify(csp, &cl.ECDSAKeyGenOpts{Curve: elliptic.P224()}, nil)

	fmt.Println("ECDSA with P256 mode test ...")
	signAndVerify(csp, &cl.ECDSAKeyGenOpts{Curve: elliptic.P256()}, nil)

	fmt.Println("ECDSA with P384 mode test ...")
	signAndVerify(csp, &cl.ECDSAKeyGenOpts{Curve: elliptic.P384()}, nil)

	fmt.Println("ECDSA with P521 mode test ...")
	signAndVerify(csp, &cl.ECDSAKeyGenOpts{Curve: elliptic.P521()}, nil)

	fmt.Println("RSA with default schema test ...")
	signAndVerify(csp, &cl.RSAKeyGenOpts{}, nil)

	fmt.Println("RSA with PSS schema test ...")
	signAndVerify(csp, &cl.RSAKeyGenOpts{Bits: 1024}, &cl.RSASignatureOpts{Schema: cl.PSS})

	fmt.Println("RSA with PKCS1V15 schema test ...")
	signAndVerify(csp, &cl.RSAKeyGenOpts{Bits: 2048}, &cl.RSASignatureOpts{Schema: cl.PKCS1V15})

	fmt.Println("SM2 with default mode test ...")
	signAndVerify(csp, &cl.SM2KeyGenOpts{}, nil)
}

func signAndVerify(csp cl.CSP, keyGenOpts cl.KeyGenOpts, signOpts cl.SignatureOpts) error {
	k, err := csp.KeyGen(keyGenOpts)
	if err != nil {
		fmt.Printf("KeyGen failed: %v\n", err)
		return err
	}

	kd, _ := k.Bytes()
	fmt.Println("Private Key: ", hex.EncodeToString(kd))

	msg := []byte("hello,world")
	digest, err := csp.Hash(msg, &cl.SHA256Opts{})
	if err != nil {
		fmt.Printf("Hash failed: %v\n", err)
		return err
	}

	signature, err := csp.Sign(k, digest, signOpts)
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

	valid, err := csp.Verify(pubKey, digest, signature, signOpts)
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
