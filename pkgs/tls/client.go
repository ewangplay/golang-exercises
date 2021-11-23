package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	// Get CA certificate
	caCertFile := "ca_cert.pem"
	caCert, err := GetCert(caCertFile)
	if err != nil {
		log.Println("Failed to get CA certificate:", err)
		return
	}

	// Get client certificate
	clientCertFile := "client_cert.pem"
	clientCert, err := GetCert(clientCertFile)
	if err != nil {
		log.Println("Failed to get client certificate:", err)
		return
	}

	// Get client private key
	clientKeyFile := "client_key.pem"
	clientKey, err := GetPrivateKey(clientKeyFile)
	if err != nil {
		log.Println("Failed to get client private key:", err)
		return
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(caCert)

	cert := tls.Certificate{
		Certificate: [][]byte{clientCert.Raw},
		PrivateKey:  clientKey,
	}

	config := tls.Config{
		// InsecureSkipVerify controls whether a client verifies the server's
		// certificate chain and host name. If InsecureSkipVerify is true, crypto/tls
		// accepts any certificate presented by the server and any host name in that
		// certificate. In this mode, TLS is susceptible to machine-in-the-middle
		// attacks unless custom verification is used. This should be used only for
		// testing or in combination with VerifyConnection or VerifyPeerCertificate.
		//InsecureSkipVerify: true,
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}

	addr := "127.0.0.1:443"
	conn, err := tls.Dial("tcp", addr, &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()

	log.Println("client: connected to: ", conn.RemoteAddr())

	data := []byte("Hello")
	//conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	n, err := conn.Write(data)
	if err != nil {
		log.Println("Client: writing data error:", err)
		return
	}

	log.Printf("client: wrote %s (%d bytes)", data, n)

	buf := make([]byte, 256)
	//conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	n, err = conn.Read(buf)
	if err != nil {
		log.Println("Client: reading data error:", err)
		return
	}

	log.Printf("client: read %s (%d bytes)", string(buf[:n]), n)

	log.Print("client: exiting")

}

func GetCert(filename string) (cert *x509.Certificate, err error) {
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

func GetPrivateKey(filename string) (key *rsa.PrivateKey, err error) {
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
