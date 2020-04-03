package crypto

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/theckman/go-securerandom"
	"log"
)

func GenerateBase58Str() string {
	rndStr := generateRnd58(12)
	return rndStr[len(rndStr)-8:]
}

func generateRnd58(n int) string {
	bytes, err := securerandom.Bytes(n)
	if err != nil {
		log.Fatal(err)
	}
	return base58.Encode(bytes)
}
