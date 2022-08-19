package util

import "math/big"

func Min(a, b *big.Float) *big.Float {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}
