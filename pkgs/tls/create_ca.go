package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	mrand "math/rand"
	"os"
	"time"
)

func main() {

	// Create CA certificate template
	template := &x509.Certificate{
		SerialNumber: big.NewInt(mrand.Int63()),
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Organization:       []string{"CIECC"},
			OrganizationalUnit: []string{"GFA"},
			Province:           []string{"Beijing"},
			CommonName:         "GFACA",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(10, 0, 0),
		// If SubjectKeyId from template is empty and the template is a CA,
		// SubjectKeyId will be generated from the hash of the public key.
		// SubjectKeyId:          []byte{1, 2, 3, 4, 5},
		BasicConstraintsValid: true,
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	// Generate CA RSA key
	priv, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Println("Failed to generate RSA key:", err)
		return
	}
	pub := &priv.PublicKey

	// Create CA certificate
	// The certificate is signed by parent. If parent is equal
	//  to template then the certificate is self-signed.
	// The parameter pub is the public key of the certificate
	//  to be generated and priv is the private key of the signer.
	derCertData, err := x509.CreateCertificate(rand.Reader, template, template, pub, priv)
	if err != nil {
		log.Println("Failed to create CA certificate:", err)
		return
	}

	// Write CA certificate to file
	certFile := "ca_cert.pem"
	err = WriteCertificateToFile(certFile, derCertData)
	if err != nil {
		log.Println("Failed to write CA certificate:", err)
		return
	}
	log.Println("Successfully write CA certificate to file:", certFile)

	// Write CA private key to file
	privFile := "ca_key.pem"
	err = WritePrivateKeyToFile(privFile, priv)
	if err != nil {
		log.Println("Failed to write CA private key:", err)
		return
	}
	log.Println("Successfully write CA private key to file:", privFile)
}

func WriteCertificateToFile(filename string, derCertData []byte) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	b := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derCertData,
	}
	return pem.Encode(f, b)
}

func WritePrivateKeyToFile(filename string, privKey *rsa.PrivateKey) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	derKeyData := x509.MarshalPKCS1PrivateKey(privKey)

	b := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derKeyData,
	}
	return pem.Encode(f, b)
}
