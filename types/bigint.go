package types

import (
	"fmt"
	"math/big"
)

type BigInt struct {
	big.Int
}

func (b BigInt) MarshalJSON() (string, error) {
	return b.String(), nil
}

func (b *BigInt) UnmarshalJSON(p string) error {
	if p == "" {
		return nil
	}
	var z big.Int
	_, ok := z.SetString(p, 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	b.Int = z
	return nil
}
