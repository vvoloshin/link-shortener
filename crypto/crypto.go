package crypto

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/theckman/go-securerandom"
)

func GenerateBase58Str() string {
	var result string
	for {
		rndStr, err := generateRnd58(12)
		if err == nil {
			result = rndStr
			break
		}
	}
	return result[len(result)-8:]
}

func generateRnd58(n int) (string, error) {
	bytes, err := securerandom.Bytes(n)
	if err != nil {
		return "", err
	}
	return base58.Encode(bytes), nil
}
