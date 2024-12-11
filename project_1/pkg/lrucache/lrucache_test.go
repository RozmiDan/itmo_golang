package lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var _ LRU[float32, int64] = (*LruCache[float32, int64])(nil)

func TestLru(t *testing.T) {
	t.Parallel()
	a := NewLruCache[int, string](5,"no")
	a.Put(1, "first")
	a.Put(2, "second")
	a.Put(3, "third")
	a.Put(4, "fouth")
	a.Put(5, "fivth")

	ar1, ar2 := getKeysNValues(a)
	require.Equal(t, ar1, []int{5,4,3,2,1})
	require.Equal(t, ar2, []string{"fivth","fouth","third","second","first"})

	a.Put(6, "sixth")
	ar1, ar2 = getKeysNValues(a)
	require.Equal(t, ar1, []int{6,5,4,3,2})
	require.Equal(t, ar2, []string{"sixth","fivth","fouth","third","second"})
	require.Equal(t, "fouth", a.Get(4))

}

func getKeysNValues[K comparable, V any] (lru *LruCache[K,V]) ([]K, []V) {
	arr_1 := make([]K, 0)
	arr_2 := make([]V, 0)

	for k, v := range lru.All(){
		arr_1 = append(arr_1, k)
		arr_2 = append(arr_2, v)
	}

	return arr_1, arr_2
}