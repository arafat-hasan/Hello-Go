package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc321!@#$%^&*()_+'`~{}|[]<>,.?;:"
	fmt.Println(data)

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	sEnc = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ = base64.URLEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	url := "https://www.prothomalo.com/opinion/column/ইলিশ-কেন-বেজার"
	fmt.Println(url)

	sEnc = base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(sEnc)

	sDec, _ = base64.URLEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

}
