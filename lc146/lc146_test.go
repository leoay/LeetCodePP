package lc146

import (
	"fmt"
	"testing"
)

func TestSol(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	fmt.Println(lruCache)
	lruCache.Put(2, 2)
	fmt.Println(lruCache.Get(1))
	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(2))
	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(3))
	fmt.Println(lruCache.Get(4))
}
