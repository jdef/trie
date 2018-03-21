package trie

import (
	"crypto/rand"
	"testing"
)

var stringKeys [1000]string // random string keys
const bytesPerKey = 30

var pathKeys [1000]string // random /paths/of/parts keys
const partsPerKey = 3     // (e.g. /a/b/c has parts /a, /b, /c)
const bytesPerPart = 10

var noValue = (interface{})(struct{}{})

func init() {
	// string keys
	for i := 0; i < len(stringKeys); i++ {
		key := make([]byte, bytesPerKey)
		if _, err := rand.Read(key); err != nil {
			panic("error generating random byte slice")
		}
		stringKeys[i] = string(key)
	}

	// path keys
	for i := 0; i < len(pathKeys); i++ {
		var key string
		for i := 0; i < partsPerKey; i++ {
			key += "/"
			part := make([]byte, bytesPerPart)
			if _, err := rand.Read(part); err != nil {
				panic("error generating random byte slice")
			}
			key += string(part)
		}
		pathKeys[i] = string(key)
	}
}

// RuneTrie
///////////////////////////////////////////////////////////////////////////////

// string keys

func BenchmarkRuneTriePutStringKey(b *testing.B) {
	trie := NewRuneTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Put(stringKeys[i%len(stringKeys)], noValue)
	}
}

func BenchmarkRuneTrieGetStringKey(b *testing.B) {
	trie := NewRuneTrie()
	for i := 0; i < b.N; i++ {
		trie.Put(stringKeys[i%len(stringKeys)], noValue)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Get(stringKeys[i%len(stringKeys)])
	}
}

// path keys

func BenchmarkRuneTriePutPathKey(b *testing.B) {
	trie := NewRuneTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Put(pathKeys[i%len(pathKeys)], noValue)
	}
}

func BenchmarkRuneTrieGetPathKey(b *testing.B) {
	trie := NewRuneTrie()
	for i := 0; i < b.N; i++ {
		trie.Put(pathKeys[i%len(pathKeys)], noValue)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Get(pathKeys[i%len(pathKeys)])
	}
}

// PathTrie
///////////////////////////////////////////////////////////////////////////////

// string keys

func BenchmarkPathTriePutStringKey(b *testing.B) {
	trie := NewPathTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Put(stringKeys[i%len(stringKeys)], noValue)
	}
}

func BenchmarkPathTrieGetStringKey(b *testing.B) {
	trie := NewPathTrie()
	for i := 0; i < b.N; i++ {
		trie.Put(stringKeys[i%len(stringKeys)], noValue)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Get(stringKeys[i%len(stringKeys)])
	}
}

// path keys

func BenchmarkPathTriePutPathKey(b *testing.B) {
	trie := NewPathTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Put(pathKeys[i%len(pathKeys)], noValue)
	}
}

func BenchmarkPathTrieGetPathKey(b *testing.B) {
	trie := NewPathTrie()
	for i := 0; i < b.N; i++ {
		trie.Put(pathKeys[i%len(pathKeys)], noValue)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Get(pathKeys[i%len(pathKeys)])
	}
}

// benchmark PathSegmenter

func BenchmarkPathSegmenter(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for j := 0; j < b.N; j++ {
		for part, i := PathSegmenter(pathKeys[j%len(pathKeys)], 0); ; part, i = PathSegmenter(pathKeys[j%len(pathKeys)], i) {
			var _ = part // NoOp 'use' the key part
			if i == -1 {
				break
			}
		}
	}
}
