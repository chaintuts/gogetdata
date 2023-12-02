// This file contains code for a simple static data web server
//
// Author: Josh McIntyre
//

package main

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"
)

// Constants
const CERT_FILE = "../cert.crt"
const KEY_FILE = "../cert.key"

const PORT = ":443"

// Handle a simple request by reading a file and displaying contents
func request_handler(writer http.ResponseWriter, request *http.Request) {

	// RequestURI gives the URI after the host, including the root slash
	// ex: http://localhost/data.html is returned as /data.html
	// So, trim the / off using a slice to get the right filename
	file_requested := request.RequestURI[1:]
	fmt.Printf("%s\n", file_requested)

	// Read the requested file
	content, err := os.ReadFile(file_requested)
	if err != nil {
		// Write a simple error message if unable to read
		fmt.Fprintf(writer, "Unable to read requested file")
		return
	}
	
	// Write the data to the ResponseWriter
	fmt.Fprintf(writer, string(content))

}

// Generate a certificate for HTTPS use
// Write the cert and key to files on disk
func generate_cert() {

	// Determine if we should generate a new certificate or if one already exists
	generate := false
	_, err := os.ReadFile(CERT_FILE)
	if err != nil {
		generate = true
	}

	_, err = os.ReadFile(CERT_FILE)
	if err != nil {
		generate = true
	}
	
	if !generate {
		fmt.Println("Certificate and private key already exist, using existing")
		return
	}
	
	// If generating a new certificate
	// Generate a random RSA private key
    key, err := rsa.GenerateKey(rand.Reader, 4096)
    if err != nil {
        fmt.Println(err)
    }

	// Generate the certificate
    cert_template := x509.Certificate {

        SerialNumber: big.NewInt(0),
        Subject: pkix.Name {CommonName: "localhost"},
        SignatureAlgorithm: x509.SHA256WithRSA,
        NotBefore: time.Now(),
        NotAfter: time.Now().Add(365*24*10*time.Hour),
    }

    cert_bytes, err := x509.CreateCertificate(rand.Reader, &cert_template, &cert_template, &key.PublicKey, key)

    if err != nil {
        fmt.Println(err)
    }

    // Encode the generated private key and certificate
    key_x509_bytes := x509.MarshalPKCS1PrivateKey(key)
    key_pem := string(pem.EncodeToMemory(&pem.Block{ 
                                              Type: "RSA PRIVATE KEY",
                                              Bytes: key_x509_bytes,
                                            },))

    cert_pem := string(pem.EncodeToMemory(&pem.Block{
                                                     Type: "CERTIFICATE",
                                                     Bytes: cert_bytes,
                                                },))
    
	fmt.Println("Generated new certificate with private key")
    fmt.Println(cert_pem)

	// Write the certificate and private key encoded data to file
    err = os.WriteFile(KEY_FILE, []byte(key_pem), 0775)
    if err != nil {
        fmt.Println(err)
    }

    err = os.WriteFile(CERT_FILE, []byte(cert_pem), 0775)
    if err != nil {
        fmt.Println(err)
    }
}

// The main entry point for the program
func main() {
	
	// Generate certificate
	generate_cert()
	
	// Define URL handlers
	http.HandleFunc("/", request_handler)

	// Start the web server
	http.ListenAndServeTLS(PORT, CERT_FILE, KEY_FILE, nil)
}