package go_diffie_hellman

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

var (
	digits              = "0123456789"
	RandBigIntParamsErr = errors.New("params n must > 0")
)

func Prime(n int) (bi *big.Int, err error) {
	bi = new(big.Int)
	for {
		bi, err = randBigInt(n)
		if err != nil {
			return
		}
		// If bi is prime, returns true.
		if bi.ProbablyPrime(20) {
			return
		}
	}
}

func randBigInt(n int) (bi *big.Int, err error) {
	bi = new(big.Int)

	if n <= 0 {
		return bi, RandBigIntParamsErr
	}

	s := make([]byte, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			idx, _ := rand.Int(rand.Reader, big.NewInt(9))
			s[i] = digits[1:][int(idx.Int64())]
			continue
		}
		idx, _ := rand.Int(rand.Reader, big.NewInt(10))
		s[i] = digits[int(idx.Int64())]
	}

	_, _ = fmt.Sscan(string(s), bi)
	return
}
