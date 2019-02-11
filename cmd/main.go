package main

import (
	"fmt"

	"github.com/dbenque/simple-cache/pkg/lru"
)

func main() {

	values := map[string]string{
		"1":  "Tiernan Lord",
		"2":  "Lawrence Lennon",
		"3":  "Harris Sheppard",
		"4":  "Yu Bannister",
		"5":  "Kaira Bone",
		"6":  "Aaran Greig",
		"7":  "Rami Baird",
		"8":  "Anisa Kidd",
		"9":  "Bruno Wilcox",
		"10": "Antony Lugo",
	}

	lruCache := lru.NewLRUCache(6)
	for k, v := range values {
		fmt.Printf("Insertion of (%s,%s)\n", k, v)
		lruCache.Set(k, v)
		lruCache.Stat()
	}
	for k, v := range values {
		fmt.Printf("Insertion of (%s,%s)\n", k, v)
		lruCache.Set(k, v)
		lruCache.Stat()
	}
}
