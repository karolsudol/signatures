package signatures

// "github.com/ethereum/go-ethereum/common/hexutil"

// PublicKeyBytesToAddress ...
// func publicKeyBytesToAddress(publicKey []byte) common.Address {
// 	var buf []byte

// 	hash := sha3.NewLegacyKeccak256()
// 	hash.Write(publicKey[1:]) // remove EC prefix 04
// 	buf = hash.Sum(nil)
// 	address := buf[12:]

// 	return common.HexToAddress(hex.EncodeToString(address))
// }

// func Verify(msg []string, signature []byte, pubKey []byte) (verified bool, err error) {
// 	verified = false

// 	data := []byte(strings.Join(msg, " "))
// 	hash := crypto.Keccak256Hash(data)

// 	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
// 	if err != nil {
// 		return verified, fmt.Errorf("crypto.Ecrecover sig err: %v", err)
// 	}

// 	publicKeyBytes, _ := hex.DecodeString(pubKey)
// 	address := publicKeyBytesToAddress(publicKeyBytes)

// 	matches := bytes.Equal(sigPublicKey, publicKeyBytes)

// 	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
// 	if err != nil {
// 		return verified, fmt.Errorf("crypto.SigToPub err: %v", err)
// 	}

// 	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
// 	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
// 	fmt.Println(matches) // true

// 	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
// 	verified = crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
// 	fmt.Println(verified)

// 	return verified, nil

// }
