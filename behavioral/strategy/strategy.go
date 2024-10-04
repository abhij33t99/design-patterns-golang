package behavioral

import "fmt"

type EvictionAlgo interface {
	evict(*Cache)
}

//concrete strategy

type FIFO struct{}

func (f *FIFO) evict(c *Cache) {
	fmt.Println("Evicting by FIFO")
}

type LRU struct{}

func (l *LRU) evict(c *Cache) {
	fmt.Println("Evicting by LRU")
}