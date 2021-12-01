package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig(),
	}

	log.Printf("Serving on https://localhost:8000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Protocol: %s", r.Proto)))
		log.Printf("Got connection: %s from %s", r.Proto, r.RemoteAddr)
	})

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}

}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("./ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile("./ca.key")
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.X509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}
}
