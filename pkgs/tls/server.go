package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	ca_b, _ := ioutil.ReadFile("ca.pem")
	ca, _ := x509.ParseCertificate(ca_b)
	priv_b, _ := ioutil.ReadFile("ca.key")
	priv, _ := x509.ParsePKCS1PrivateKey(priv_b)

	pool := x509.NewCertPool()
	pool.AddCert(ca)

	cert := tls.Certificate{
		Certificate: [][]byte{ca_b},
		PrivateKey:  priv,
	}

	config := tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    pool,
	}

	config.Rand = rand.Reader
	service := "0.0.0.0:443"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Print("server: listening")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}

		log.Printf("server: accepted from %s", conn.RemoteAddr())
		tlsConn, ok := conn.(*tls.Conn)
		if ok {
			log.Print("ok=true")
			state := tlsConn.ConnectionState()
			for _, v := range state.PeerCertificates {
				log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
			}
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) error {
	defer conn.Close()

	data, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Printf("read data from client error: %v", err)
		return err
	}

	log.Printf("recv data from client: %s", data)

	reply := "World\n"
	_, err = io.WriteString(conn, reply)
	if err != nil {
		log.Printf("write data to client error: %v", err)
		return err
	}

	log.Printf("reply data to client succ")

	return nil
}
