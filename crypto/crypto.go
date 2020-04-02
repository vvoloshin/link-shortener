package crypto

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/theckman/go-securerandom"
	"log"
)

func GenerateBase58Str() string {
	rndStr := generateRnd58(10)
	return cutStringToLimit(rndStr, 8)
}

func generateRnd58(n int) string {
	bytes, _ := securerandom.Bytes(n)
	return base58.Encode(bytes)
}

func cutStringToLimit(s string, limit int) string {
	var res []rune
	count := 0
	for _, char := range s {
		if count == limit {
			break
		}
		res = append(res, char)
		count++
	}
	resShort := string(res)
	log.Println("debug: cutStringToLimit-string: ", resShort)
	return resShort
}
