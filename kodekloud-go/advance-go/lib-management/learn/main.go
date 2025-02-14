package main

import (
	"fmt"

	"github.com/Priyanka-yadavv/cryptit/encrypt"
	"github.com/pborman/uuid"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)

	encryptedStr := encrypt.Nimbus("kodekloud")
	fmt.Println(encryptedStr)
}
