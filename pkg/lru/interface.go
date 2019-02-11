package lru

type LRU interface {
	Get(key string) (interface{},bool)
	Set(key string, value interface{}) bool
}
