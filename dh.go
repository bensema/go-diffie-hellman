package go_diffie_hellman

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

type Dh struct {
	P *big.Int // prime number
	G *big.Int // generator of p

	AnswerKey  *big.Int
	PrivateKey *big.Int
}

func NewDh() *Dh {
	return &Dh{}
}

func (dh *Dh) Public() *big.Int {
	// g^privateKey mod p  ==> return as other side answer key
	return big.NewInt(0).Exp(dh.G, dh.PrivateKey, dh.P)
}

func (dh *Dh) Computes() *big.Int {
	// answerKey^privateKey mod p ==>this is share secret key
	return big.NewInt(0).Exp(dh.AnswerKey, dh.PrivateKey, dh.P)
}

func (dh *Dh) Sha256() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dh.Computes().String())))
}
