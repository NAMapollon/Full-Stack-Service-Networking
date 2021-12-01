package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
)

// Parse command line parameters
var httpVersion = flag.Int("version", 2, "HTTP version")
var httpAddr = flag.String("addr", "localhost", "HTTP server address")
var httpPort = flag.String("port", "8000", "HTTP server port")
var httpCert = flag.String("cert", "server.crt", "HTTP server certificate file name")

func main() {
	// Configure a parsed parameters and a http server url
	flag.Parse()
	url := fmt.Sprintf("https://%s:%s", *httpAddr, *httpPort)

	// Create a http client
	client := &http.Client{}

	// Create a pool with the server certificate since it is not signed by a known CA
	caCert, err := ioutil.ReadFile(*httpCert)
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Use the proper transport in the client
	switch *httpVersion {
	case 1:
		fmt.Printf("Connect to %s over TLS using HTTP/1.1\n", url)
		client.Transport = &http.Transport{TLSClientConfig: tlsConfig}
	case 2:
		fmt.Printf("Connect to %s over TLS using HTTP/2\n", url)
		client.Transport = &http2.Transport{TLSClientConfig: tlsConfig}
	}

	// Perform the request
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	fmt.Printf("Got response %d: %s %s\n", resp.StatusCode, resp.Proto, string(body))
}
