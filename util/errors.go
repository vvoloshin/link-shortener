package util

import "log"

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CheckErrorVerb(e error, s string) {
	if e != nil {
		log.Println(s)
		log.Fatal(e)
	}
}
