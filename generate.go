package signatures

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

func GenerateEDSA(msg []string) (signature []byte, err error) {
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

func GenerateEIP715(msg []string) (signature []byte, err error) {
	// err = godotenv.Load(".env")
	// if err != nil {
	// 	return signature, fmt.Errorf("load env key err: %v", err)
	// }

	// key := os.Getenv("PRIVATE_KEY")
	// privateKey, err := crypto.HexToECDSA(key)
	// if err != nil {
	// 	return signature, fmt.Errorf("os load env key err: %v", err)
	// }

	// var typesStandard = apitypes.Types{
	// 	"EIP712Domain": {
	// 		{
	// 			Name: "name",
	// 			Type: "string",
	// 		},
	// 		{
	// 			Name: "version",
	// 			Type: "string",
	// 		},
	// 		{
	// 			Name: "chainId",
	// 			Type: "uint256",
	// 		},
	// 		{
	// 			Name: "verifyingContract",
	// 			Type: "address",
	// 		},
	// 	},
	// 	"Mint": {
	// 		{
	// 			Name: "minter",
	// 			Type: "address",
	// 		},
	// 		{
	// 			Name: "amount",
	// 			Type: "uint256",
	// 		},
	// 	},
	// }

	// const primaryType = "Mint"

	// var domainStandard = apitypes.TypedDataDomain{
	// 	Name:              "Airdrop",
	// 	Version:           "1",
	// 	ChainId:           math.NewHexOrDecimal256(5),
	// 	VerifyingContract: "0xb5e79d544Dc8dE31991Ac395027E62B3dD43C028",
	// }

	// var messageStandard = map[string]interface{}{
	// 	"Mail": map[string]interface{}{
	// 		"name":   "Cow",
	// 		"wallet": "0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826",
	// 	},
	// 	"to": map[string]interface{}{
	// 		"name":   "Bob",
	// 		"wallet": "0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB",
	// 	},
	// 	"contents": "Hello, Bob!",
	// }

	// typedDataHash, _ := core.HashStruct(signerData.PrimaryType, signerData.Message)
	// domainSeparator, _ := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())

	// rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	// challengeHash := crypto.Keccak256Hash(rawData)

	// var typedData = apitypes.TypedData{
	// 	Types:       typesStandard,
	// 	PrimaryType: primaryType,
	// 	Domain:      domainStandard,
	// 	Message:     messageStandard,
	// }

	// ui := &headlessUi{make(chan string, 20), make(chan string, 20)}
	// am := core.StartClefAccountManager(tmpDirName(t), true, true, "")
	// api := core.NewSignerAPI(am, 51, true, ui, db, true, &storage.NoStorage{})

	// signature, err = api.SignData(context.Background(), apitypes.TextPlain.Mime, a, hexutil.Encode([]byte("EHLO world")))

	// data := []byte(strings.Join(msg, " "))
	// hash := crypto.Keccak256Hash(data)

	// signature, err = crypto.Sign(hash.Bytes(), privateKey)
	// if err != nil {
	// 	return signature, fmt.Errorf("os load env key err: %v", err)
	// }

	// fmt.Println(hexutil.Encode(signature))

	return signature, nil

}
