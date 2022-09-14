package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//createKs()
	importKs()
}

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs() {
	// keystore file
	file := "./tmp/UTC--2022-09-08T06-57-12.754954000Z--bdfc113335aca302b5d463c55277797ed6e8670a"

	// generate a new keystore object
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	// read the keystore file
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("readfile error")
		log.Fatal(err)
	}

	password := "secret"
	// import the account AND you will need to set a "new" password which in this case is the same old password
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		fmt.Println("account error")
		log.Fatal(err)
	}

	// get the Key object
	key, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatal(err)
	}

	// get private key from Key object
	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:])

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		fmt.Println("remove error")
		log.Fatal(err)
	}
}
