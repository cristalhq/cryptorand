package cryptorand

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	mathrand "math/rand"
)

// Source returns [math/rand.Source64] interface based on [crypto/rand.Reader].
// Seed method is not supported and will panic.
func Source() mathrand.Source64 {
	return newSource(rand.Reader)
}

func newSource(r io.Reader) mathrand.Source64 {
	return source{r: r}
}

type source struct {
	r io.Reader
}

func (s source) Int63() int64 {
	i, err := rand.Int(s.r, maxInt63)
	if err != nil {
		panic(fmt.Errorf("cryptorand: crypto/rand.Int returned error: %w", err))
	}
	return i.Int64()
}

func (s source) Uint64() uint64 {
	i, err := rand.Int(s.r, maxUint64)
	if err != nil {
		panic(fmt.Errorf("cryptorand: crypto/rand.Int returned error: %w", err))
	}
	return i.Uint64()
}

func (source) Seed(int64) {
	panic("cryptorand: Seed is not allowed for rand.Source")
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
