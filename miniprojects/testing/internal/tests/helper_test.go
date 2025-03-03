package tests

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/RozmiDan/miniprojects/testing/internal/cache"
	"github.com/stretchr/testify/assert"
)

func emulateLoad(t *testing.T, c cache.Cache, parallelFactor int64) {
	wg := &sync.WaitGroup{}

	for i := 0; i < int(parallelFactor); i++ {
		wg.Add(1)

		key := fmt.Sprintf("key %d", i)
		value := fmt.Sprintf("value %d", i)

		go func(key, value string) {
			defer wg.Done()
			err := c.Set(key, value)
			assert.NoError(t, err)
		}(key, value)

		wg.Add(1)

		go func(key, value string){
			defer wg.Done()
			res, err := c.Get(key)
			if !errors.Is(err, cache.ErrNotFound) {
				assert.Equal(t, value, res)
			}
		}(key, value)

		wg.Add(1)

		go func(key, value string){
			defer wg.Done()
			err := c.Delete(key)
			assert.NoError(t, err)
		}(key, value)
	}

	wg.Wait()
}