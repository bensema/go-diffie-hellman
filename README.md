# Diffie–Hellman key exchange


```go

package main

import (
	"fmt"
	dh "github.com/bensema/go-diffie-hellman"
	"math/big"
)

func main() {

	p, _ := dh.Prime(512)
	g := big.NewInt(5)

	alice := dh.NewDh()
	alice.P = p
	alice.G = g
	alice.PrivateKey, _ = dh.Prime(512)

	bob := dh.NewDh()
	bob.P = p
	bob.G = g
	bob.PrivateKey, _ = dh.Prime(512)

	alice.AnswerKey = bob.Public()
	bob.AnswerKey = alice.Public()

	fmt.Println("P:", p)
	fmt.Println("G:", g)
	fmt.Println("alice PrivateKey:", alice.PrivateKey)
	fmt.Println("bob   PrivateKey:", bob.PrivateKey)

	fmt.Println("alice AnswerKey:", alice.AnswerKey)
	fmt.Println("bob   AnswerKey:", bob.AnswerKey)

	fmt.Println("alice computes:", alice.Computes())
	fmt.Println("bob   computes:", bob.Computes())

	fmt.Println("alice sha256:", alice.Sha256())
	fmt.Println("bob   sha256:", bob.Sha256())

}

```
