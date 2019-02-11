package lru

import "fmt"

type LRUCache struct {
	nodesMap map[string]*node
	list     *linkedList
	doneChan chan struct{}
}

func NewLRUCache(maxSize int) *LRUCache {
	l := &LRUCache{
		nodesMap: map[string]*node{},
		list:     NewList(maxSize),
	}

	go func() {
		for {
			select {
			case <-l.doneChan:
				return
			case key := <-l.list.deletedKey:
				delete(l.nodesMap, key)
			}
		}
	}()
	return l
}

func (lru *LRUCache) Get(key string) (interface{}, bool) {
	if node, found := lru.nodesMap[key]; found {
		return node.value, true
	}
	return nil, false
}
func (lru *LRUCache) Set(key string, value interface{}) bool {
	if node, found := lru.nodesMap[key]; found {
		node.value = value
		lru.list.MoveToTail(node)
		return true
	}
	lru.nodesMap[key] = lru.list.Append(key, value)
	return false
}
func (lru *LRUCache) Stat() {
	fmt.Printf("Count: %d\n", len(lru.nodesMap))
	fmt.Printf("KeySequence: ")
	currentNode := lru.list.head
	for currentNode != nil {
		fmt.Printf("%s ", currentNode.key)
		currentNode = currentNode.next
	}
	fmt.Printf("\n\n")
}
