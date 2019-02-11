package lru

type node struct {
	previous, next *node
	key            string
	value          interface{}
}

type linkedList struct {
	head       *node
	tail       *node
	size       int
	maxSize    int
	deletedKey chan string
}

func NewList(maxSize int) *linkedList {
	return &linkedList{
		maxSize:    maxSize,
		deletedKey: make(chan string, 10),
	}
}

func (ll *linkedList) Append(key string, value interface{}) *node {
	defer func() {
		ll.size++
		for ll.size > ll.maxSize {
			ll.deletedKey <- ll.head.key
			ll.RemoveHead()
		}
	}()

	if ll.size == 0 {
		newNode := &node{next: ll.head, key: key, value: value}
		ll.head = newNode
		ll.tail = newNode
		return newNode
	}

	newNode := &node{previous: ll.tail, key: key, value: value}
	ll.tail.next = newNode
	ll.tail = newNode
	return newNode
}

func (ll *linkedList) MoveToTail(n *node) {
	if n.next == nil { // already tail
		return
	}
	if n.previous != nil {
		n.previous.next = n.next
	} else { // n is currently head!
		ll.head = n.next
	}
	n.next.previous = n.previous

	ll.tail.next = n
	n.previous = ll.tail
	ll.tail = n
}

func (ll *linkedList) RemoveHead() {
	if ll.head == nil {
		return
	}
	newHead := ll.head.next
	newHead.previous = nil
	ll.head.next = nil
	ll.head = newHead
	ll.size--
}
