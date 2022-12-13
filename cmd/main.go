package main

import (
	"fmt"
	"log"

	"github.com/karolsudol/signatures"
)

func main() {

	sig, err := signatures.GenerateEDSA([]string{"0x1f190F523deBD185183d8Afe76e4587a08bb84e7", "2"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(sig)

}
