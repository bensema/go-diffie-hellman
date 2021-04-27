package go_diffie_hellman

import (
	"testing"
)

func TestPrime(t *testing.T) {
	t.Parallel()
	var pairs = []struct {
		err error
		n   int
	}{
		{RandBigIntParamsErr, -1},
		{RandBigIntParamsErr, 0},
		{nil, 1},
		{nil, 2},
		{nil, 3},
		{nil, 4},
		{nil, 5},
		{nil, 6},
		{nil, 7},
	}
	for _, tt := range pairs {
		prime, err := Prime(tt.n)
		if err != tt.err {
			t.Errorf("prime: %s error: %s", prime, err)
		}
		//fmt.Println("prime:", prime)
	}
}
