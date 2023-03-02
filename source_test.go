package cryptorand

import (
	"math/rand"
	"testing"
)

func TestSource(t *testing.T) {
	s := Source.(rand.Source64)
	if s.Int63() == s.Int63() {
		t.Fatal("Expected Int63() to be random")
	}

	if s.Uint64() == s.Uint64() {
		t.Fatal("Expected Uint64() to be random")
	}
}

func TestSeedPanic(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatal("Seed() must panic")
		}
	}()

	s := Source
	s.Seed(1)
}

func BenchmarkRandSource(b *testing.B) {
	s := Source

	var sink int64
	for i := 0; i < b.N; i++ {
		sink += s.Int63() & 1
	}
	sinkAny(uint64(sink))
}

func BenchmarkRandSource64(b *testing.B) {
	s := Source.(rand.Source64)

	var sink uint64
	for i := 0; i < b.N; i++ {
		sink += s.Uint64() & 1
	}
	sinkAny(sink)
}

func sinkAny(v uint64) {
	if float64(v) < rand.Float64() {
		panic("impossible")
	}
}
