package counter_test

import (
	"math/rand"
	"testing"

	"github.com/benhoyt/counter"
)

const (
	numWords = 10000
	wordLen  = 7
)

func BenchmarkMostlyUniqueCounter(b *testing.B) {
	words := make([][]byte, numWords)
	for i := 0; i < numWords; i++ {
		words[i] = randomChars(wordLen)
	}
	var counts counter.Counter
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords; wi++ {
			counts.Inc(words[wi], 1)
		}
	}
}

func BenchmarkNonUniqueCounter(b *testing.B) {
	words := make([][]byte, numWords/10)
	for i := 0; i < numWords/10; i++ {
		words[i] = randomChars(wordLen)
	}
	var counts counter.Counter
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords/10; wi++ {
			counts.Inc(words[wi], 1)
			counts.Inc(words[wi], 1)
			counts.Inc(words[wi], 1)
			counts.Inc(words[wi], 1)
			counts.Inc(words[wi], 1)
			counts.Inc(words[wi], 1)
			counts.Inc(words[wi], 1)
			counts.Inc(words[wi], 1)
			counts.Inc(words[wi], 1)
			counts.Inc(words[wi], 1)
		}
	}
}

func BenchmarkMostlyUniqueMapBytes(b *testing.B) {
	words := make([][]byte, numWords)
	for i := 0; i < numWords; i++ {
		words[i] = randomChars(wordLen)
	}
	counts := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords; wi++ {
			counts[string(words[wi])]++
		}
	}
}

func BenchmarkNonUniqueMapBytes(b *testing.B) {
	words := make([][]byte, numWords/10)
	for i := 0; i < numWords/10; i++ {
		words[i] = randomChars(wordLen)
	}
	counts := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords/10; wi++ {
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
			counts[string(words[wi])]++
		}
	}
}

func BenchmarkMostlyUniqueMapString(b *testing.B) {
	words := make([]string, numWords)
	for i := 0; i < numWords; i++ {
		words[i] = string(randomChars(wordLen))
	}
	counts := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords; wi++ {
			counts[words[wi]]++
		}
	}
}

func BenchmarkNonUniqueMapString(b *testing.B) {
	words := make([]string, numWords/10)
	for i := 0; i < numWords/10; i++ {
		words[i] = string(randomChars(wordLen))
	}
	counts := make(map[string]int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for wi := 0; wi < numWords/10; wi++ {
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
			counts[words[wi]]++
		}
	}
}

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func randomChars(length int) []byte {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return b
}
