package linklist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinklist(t *testing.T) {
	t.Parallel()

	b := NewList[string]()

	b.PushBack("1")
	b.PushFront("2")
	b.PushBack("3")
	b.PushFront("4")
	b.PushBack("5")
	b.PushFront("6")
	b.PushBack("7")
	b.PushFront("8")

	require.Equal(t, "8", b.Front().data)
	require.Equal(t, "7", b.Back().data)


	a := NewList[string]()
	a.PushBack("1")
	a.Remove(a.Back())
	require.Equal(t, a.Size(), 0)
	a.PushBack("2")
	a.Remove(a.Front())
	require.Equal(t, a.Size(), 0)

	a.PushBack("3")
	a.PushFront("4")
	a.PushBack("5")
	a.PushFront("6")
	a.PushBack("7")
	a.PushBack("8")
	a.Remove(a.Front())
	a.Remove(a.Back())
	a.Remove(a.Front())
	require.Equal(t, "3", a.Front().data)
	require.Equal(t, "7", a.Back().data)
	// 3 5 7

	arr := make([]string, 0)
	
	for it := range a.All(){
		arr = append(arr, it)
	}
	require.Equal(t, arr, []string{"3", "5", "7"})
}


