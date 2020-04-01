package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/big"
)

func Encode62(payload string) string {
	hexStr := toHmac(payload)
	return to62(hexStr)
}

//кодирование при помощи ключа в 16-ричное хэш-значение
func toHmac(s string) string {
	secret := "secret"
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(s))
	hexStr := hex.EncodeToString(h.Sum(nil))
	log.Println("debug: hex-string: ", hexStr)
	return hexStr
}

//кодирование в 62-ричную систему
func to62(s string) string {
	bigInt := new(big.Int)
	bigInt.SetString(s, 16)
	text62 := bigInt.Text(62)
	log.Println("debug: 62-digit-string: ", text62)
	return text62
}
