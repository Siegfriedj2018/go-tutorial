package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	// "github.com/golang-jwt/jwt/v5"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Choose a signing method:")
	fmt.Println("1. RSA")
	fmt.Println("2. ECDSA")

	for {
		fmt.Print("Enter your choice (1 or 2): ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		switch choiceStr {
		case "1":
			generateRSAKeys()
			return
		case "2":
			generateECDSAKeys()
			return
		default:
			fmt.Println("Invalid choice. Please enter 1 or 2.")
		}
	}
}

func generateRSAKeys() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter desired bits for RSA key (e.g., 2048): ")
	var bits int
	_, err := fmt.Fscan(reader, &bits)
	if err != nil {
		log.Fatal("Failed to read input:", err)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Fatal("Failed to generate private key:", err)
	}

	saveKeyToFile("private_key.pem", privateKey)

	publicKey := &privateKey.PublicKey
	saveKeyToFile("public_key.pem", publicKey)

	fmt.Println("RSA keys generated successfully!")
}

func generateECDSAKeys() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		log.Fatal("Failed to generate private key:", err)
	}

	saveKeyToFile("private_key.pem", privateKey)

	publicKey := &privateKey.PublicKey
	saveKeyToFile("public_key.pem", publicKey)

	fmt.Println("ECDSA keys generated successfully!")
}

func saveKeyToFile(filename string, key interface{}) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the directory to save the keys: ")
	directory, _ := reader.ReadString('\n')
	directory = strings.TrimSpace(directory)

	var pemBlock *pem.Block
	switch k := key.(type) {
	case *rsa.PrivateKey:
		pemBlock = &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(k),
		}
	case *rsa.PublicKey:
		pemBlock = &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(k),
		}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			log.Fatal("Failed to marshal ECDSA private key:", err)
		}
		pemBlock = &pem.Block{Type: "ECDSA PRIVATE KEY", Bytes: b}
	case *ecdsa.PublicKey:
		b, err := x509.MarshalPKIXPublicKey(k)
		if err != nil {
			log.Fatal("Failed to marshal ECDSA public key:", err)
		}
		pemBlock = &pem.Block{Type: "ECDSA PUBLIC KEY", Bytes: b}
	default:
		log.Fatalf("Unsupported key type %T", k)
	}

	pemBytes := pem.EncodeToMemory(pemBlock)

	filePath := fmt.Sprintf("%s/%s", directory, filename)
	err := ioutil.WriteFile(filePath, pemBytes, 0644)
	if err != nil {
		log.Fatalf("Failed to write key to file: %v", err)
	}

	fmt.Printf("%s saved to %s\n", filename, filePath)
}