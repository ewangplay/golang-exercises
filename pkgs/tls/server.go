package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	// Get CA certificate
	caCertFile := "ca_cert.pem"
	caCert, err := GetCert(caCertFile)
	if err != nil {
		log.Println("Failed to get CA certificate:", err)
		return
	}
	// Get server certificate
	serverCertFile := "server_cert.pem"
	serverCert, err := GetCert(serverCertFile)
	if err != nil {
		log.Println("Failed to get server certificate:", err)
		return
	}

	// Get server private key
	serverKeyFile := "server_key.pem"
	serverKey, err := GetPrivateKey(serverKeyFile)
	if err != nil {
		log.Println("Failed to get server private key:", err)
		return
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(caCert)

	cert := tls.Certificate{
		Certificate: [][]byte{serverCert.Raw},
		PrivateKey:  serverKey,
	}

	config := tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	}

	addr := "0.0.0.0:443"
	listener, err := tls.Listen("tcp", addr, &config)
	if err != nil {
		log.Println("server listening failed:", err)
	}
	log.Println("server listening on", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}

		log.Printf("server: accepted from %s", conn.RemoteAddr())

		/*
			tlsConn, ok := conn.(*tls.Conn)
			if ok {
				log.Print("ok=true")
				state := tlsConn.ConnectionState()
				for _, v := range state.PeerCertificates {
					log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
				}
			}
		*/

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) error {
	defer conn.Close()

	buf := make([]byte, 256)
	//conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("server: reading data error:", err)
		return err
	}

	log.Printf("server: recv data from client: %s (%d bytes)", buf[:n], n)

	reply := []byte("World")
	//conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
	n, err = conn.Write(reply)
	if err != nil {
		log.Println("Client: writing data error:", err)
		return err
	}

	log.Printf("server: reply data to client succ: %s (%d bytes)", reply, n)

	return nil
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
