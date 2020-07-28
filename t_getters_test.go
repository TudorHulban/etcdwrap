package etcdwrap

import (
	"context"
	"testing"

	"github.com/TudorHulban/loginfo"
	"github.com/stretchr/testify/assert"
)

// Target of test:
// a. insert with prefix works
func TestSimpleGet(t *testing.T) {
	l, errLog := loginfo.New(2)
	assert.Nil(t, errLog)

	storeClient, err := NewETCDClient([]string{etcdURL}, l)
	assert.Nil(t, err)
	assert.NotNil(t, storeClient)

	defer func() {
		assert.Nil(t, storeClient.TheStore.Close())
	}()

	// a.
	assert.Nil(t, storeClient.SetWPrefixKV(context.Background(), testKeyPrefix, KV{
		key:   testKey,
		value: testValue1,
	}))
	assert.Nil(t, storeClient.SetWPrefixKV(context.Background(), testKeyPrefix, KV{
		key:   testKey,
		value: testValue2,
	}))

	// b.
	sliceKV, errGetByPrefix := storeClient.GetKVByKPrefix(context.Background(), testKeyPrefix)
	assert.Nil(t, errGetByPrefix)
	assert.Equal(t, 0, len(sliceKV))
}
