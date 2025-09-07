package cryptorand_test

import (
	"fmt"
	"math/rand"

	"github.com/cristalhq/cryptorand"
)

func Example() {
	r := rand.New(cryptorand.Source())

	m := map[string]struct{}{}
	for i := 0; i < 100; i++ {
		s := fmt.Sprint(r.Float64())
		m[s] = struct{}{}
	}

	fmt.Printf("Have %d unique floats", len(m))

	// Output:
	// Have 100 unique floats
}
