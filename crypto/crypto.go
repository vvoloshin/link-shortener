package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/big"
)

func Encode(payload string) string {
	hexStr := toHmac(payload)
	digit62 := to62(hexStr)
	return cutStringTo8(digit62)
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

//укорачивание строки до 6 символов, пропуск некоторых двузначных символов
func cutStringTo8(s string) string {
	excl := []rune{'0', 'O', 'l', 'i', 'I'}
	var res []rune
	count := 0
	for _, char := range s {
		if count == 8 {
			break
		}
		if contains(excl, char) {
			continue
		}
		res = append(res, char)
		count++
	}
	resShort := string(res)
	log.Println("debug: cutStringTo8-string: ", resShort)
	return resShort
}

//поиск руны в слайсе, true - если содержит
func contains(src []rune, r rune) bool {
	for _, n := range src {
		if r == n {
			return true
		}
	}
	return false
}
