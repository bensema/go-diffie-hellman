package go_diffie_hellman

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestDh(t *testing.T) {
	t.Parallel()
	p, _ := Prime(512)
	g := big.NewInt(5)

	alice := NewDh()
	bob := NewDh()

	// publicly agree to use a modulus p and g
	alice.P = p
	alice.G = g
	bob.P = p
	bob.G = g

	// chooses a private key
	alice.PrivateKey, _ = Prime(512)
	bob.PrivateKey, _ = Prime(512)

	//  alice public key as bob answer key
	alicePublicKey := alice.Public()
	bobPublicKey := bob.Public()

	alice.AnswerKey = bobPublicKey
	bob.AnswerKey = alicePublicKey

	assert.Equal(t, alice.Computes(), bob.Computes(), "they should be equal")
	//fmt.Println("alice computes:", alice.Computes())
	//fmt.Println("bob   computes:", bob.Computes())
	fmt.Println("alice sha256:", alice.Sha256())
	fmt.Println("bob   sha256:", bob.Sha256())
	fmt.Println("alice public:", alice.Public())
	fmt.Println("bob   public:", bob.Public())

}
