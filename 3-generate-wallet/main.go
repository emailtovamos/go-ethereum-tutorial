package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 1. Generate private key, used to sign transaction securely
	privateKey, err := crypto.GenerateKey() // 64 hex characters,
	if err != nil {
		log.Fatal(err)
	}

	// 1.1 Private key in readable format
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:]) //strip off the 0x, e.g. b8caf1fcb42cad076ba79ab5c27d58f030a18a3b85423aab28513b14f9edfa01

	// 2. Generate public key from private key, used to verify transaction
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // convert to hex
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)           // convert to bytes
	fmt.Println("public key: ", hexutil.Encode(publicKeyBytes)[4:]) //strip off the 0x and the first 2 characters 04

	// 3. Address
	// Keccak-256 hash of the public key, and then we take the last 40 characters (20 bytes) and prefix it with 0x
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)

}
