package trie

// Trier exposes the Trie structure capabilities.
type Trier interface {
	Get(key string) (interface{}, bool)
	Put(key string, value interface{}) bool
	Delete(key string) bool
	Walk(walker WalkFunc)
}
