package tests

import (
	"testing"

	cache "github.com/RozmiDan/miniprojects/testing/internal/cache"

	"github.com/stretchr/testify/assert"
)

func Test_Cache(t *testing.T) {
	t.Parallel()

	t.Run("Correctly storred value", func(t *testing.T) {
		t.Parallel()
		curInst := cache.NewCache()
		anotherInst := cache.NewCache()

		curInst.Set("first", "first_value")
		curInst.Set("second", "second_value")
		curInst.Set("third", "third_value")
		curInst.Set("fourth", "fourth_value")
		curInst.Set("fivth", "fivth_value")
		curInst.Set("sixth", "sixth_value")

		curInst.Delete("third")
		curInst.Delete("sixth")

		anotherInst.Set("first", "first_value")
		anotherInst.Set("second", "second_value")
		anotherInst.Set("fourth", "fourth_value")
		anotherInst.Set("fivth", "fivth_value")

		assert.Equal(t, anotherInst, curInst)
	})
	
	t.Run("No data races", func(t *testing.T){
		t.Parallel()

		parallelFactor := int64(100_000)
		emulateLoad(t, cache.NewCache(), parallelFactor)
	})
}