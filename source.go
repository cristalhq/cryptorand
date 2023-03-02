package cryptorand

import (
	cryptorand "crypto/rand"
	"fmt"
	"io"
	"math/big"
	"math/rand"
)

// Source implements math/rand.Source64 interface.
// Seed method is not supported and will panic.
var Source = newSource(cryptorand.Reader)

func newSource(r io.Reader) rand.Source {
	return source{r: r}
}

type source struct {
	r io.Reader
}

func (s source) Int63() int64 {
	i, err := cryptorand.Int(s.r, maxInt63)
	if err != nil {
		panic(fmt.Errorf("cryptorand: crypto/rand.Int returned error: %w", err))
	}
	return i.Int64()
}

func (s source) Uint64() uint64 {
	i, err := cryptorand.Int(s.r, maxUint64)
	if err != nil {
		panic(fmt.Errorf("cryptorand: crypto/rand.Int returned error: %w", err))
	}
	return i.Uint64()
}

func (source) Seed(int64) {
	panic("cryptorand: Seed is not allowed for cryptorand.Source")
}

var maxInt63, maxUint64 *big.Int

func init() {
	var ok bool

	// Same as '1 << 63'
	maxInt63, ok = new(big.Int).SetString("9223372036854775808", 10)
	if !ok {
		panic("cryptorand: BUG, cannot set maxInt63")
	}

	// Same as '1 << 64'
	maxUint64, ok = new(big.Int).SetString("18446744073709551616", 10)
	if !ok {
		panic("cryptorand: BUG, cannot set maxUint64")
	}
}
