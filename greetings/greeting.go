package greetings

import (
	"fmt"

	pkcs11 "github.com/miekg/pkcs11"
)

func GetGreetings() int {
	fmt.Println("greeting - hello world")
	var p *pkcs11.Ctx
	fmt.Println("ctx - ", p)
	return 111
}
