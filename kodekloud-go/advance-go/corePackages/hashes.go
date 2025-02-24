package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func encodeWithMD5(str string) string {
	var hashCode = md5.Sum([]byte(str))
	return hex.EncodeToString(hashCode[:])
}
func main() {
	var password string
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)
	fmt.Println("Password encrypted to: ", encodeWithMD5(password))
}
