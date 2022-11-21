package signatures

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

func Generate(msg []string) (signature []byte, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return signature, fmt.Errorf("load env key err: %v", err)
	}

	key := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return signature, fmt.Errorf("os load env key err: %v", err)
	}

	data := []byte(strings.Join(msg, " "))
	hash := crypto.Keccak256Hash(data)

	signature, err = crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return signature, fmt.Errorf("os load env key err: %v", err)
	}

	fmt.Println(hexutil.Encode(signature))

	return signature, nil

}
