package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	mrand "math/rand"
	"net"
	"os"
	"time"
)

func main() {

	// Parse command line arguments
	pCN := flag.String("cn", "entity", "common name")
	flag.Parse()

	// Get CA certificate
	caCertFile := "ca_cert.pem"
	caCert, err := GetCACert(caCertFile)
	if err != nil {
		log.Println("Failed to get CA certificate:", err)
		return
	}

	// Get CA private key
	caKeyFile := "ca_key.pem"
	caKey, err := GetCAPrivateKey(caKeyFile)
	if err != nil {
		log.Println("Failed to get CA private key:", err)
		return
	}

	// Create Entity certificate template
	template := &x509.Certificate{
		SerialNumber: big.NewInt(mrand.Int63()),
		Subject: pkix.Name{
			Country:            []string{"CN"},
			Organization:       []string{"CIECC"},
			OrganizationalUnit: []string{"GFA"},
			Province:           []string{"Beijing"},
			CommonName:         *pCN,
		},
		IPAddresses: []net.IP{net.IPv4(127, 0, 0, 1)},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(10, 0, 0),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	// Generate Entity RSA key
	priv, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Println("Failed to generate RSA key:", err)
		return
	}
	pub := &priv.PublicKey

	// Create Entity certificate
	derCertData, err := x509.CreateCertificate(rand.Reader, template, caCert, pub, caKey)
	if err != nil {
		log.Println("Failed to create entity certificate:", err)
		return
	}

	// Write Entity certificate to file
	certFile := fmt.Sprintf("%s_cert.pem", *pCN)
	err = WriteCertificateToFile(certFile, derCertData)
	if err != nil {
		log.Println("Failed to write entity certificate:", err)
		return
	}
	log.Println("Successfully write entity certificate to file:", certFile)

	// Write Entity private key to file
	privFile := fmt.Sprintf("%s_key.pem", *pCN)
	err = WritePrivateKeyToFile(privFile, priv)
	if err != nil {
		log.Println("Failed to write entity private key:", err)
		return
	}
	log.Println("Successfully write entity private key to file:", privFile)
}

func GetCACert(filename string) (cert *x509.Certificate, err error) {
	pemData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, rest := pem.Decode(pemData)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("failed to decode PEM block")
	}
	log.Println("the remainting data:", rest)

	cert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return
}

func GetCAPrivateKey(filename string) (key *rsa.PrivateKey, err error) {
	pemData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, rest := pem.Decode(pemData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block")
	}
	log.Println("the remainting data:", rest)

	key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return
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
