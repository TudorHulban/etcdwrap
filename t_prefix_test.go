package etcdwrap

import (
	"context"
	"testing"

	"github.com/TudorHulban/loginfo"
	"github.com/stretchr/testify/assert"
)

// Target of test:
// a. insert with prefix
// b. get by prefix
// c. get by range
// d. delete keys by prefix
// e. delete keys by range
func TestMultiple(t *testing.T) {
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
		key:   testKey1,
		value: testValue1,
	}))
	assert.Nil(t, storeClient.SetWPrefixKV(context.Background(), testKeyPrefix, KV{
		key:   testKey2,
		value: testValue2,
	}))

	// b.
	sliceKV1, errGetByPrefix := storeClient.GetKVByKPrefix(context.Background(), testKeyPrefix)
	assert.Nil(t, errGetByPrefix)
	assert.Equal(t, 2, len(sliceKV1))

	// c.
	sliceKV2, errGetByRange1 := storeClient.GetKVByKRangeFrom(context.Background(), testKeyPrefix+testKey1)
	assert.Nil(t, errGetByRange1)
	assert.Equal(t, 2, len(sliceKV2))

	// d.
	assert.Nil(t, storeClient.DeleteKVByPrefix(context.Background(), testKeyPrefix))

	sliceKV3, errGetByRange2 := storeClient.GetKVByKRangeFrom(context.Background(), testKeyPrefix+testKey1)
	assert.Error(t, errGetByRange2)
	assert.Equal(t, 0, len(sliceKV3))

	// e.
	assert.Nil(t, storeClient.SetWPrefixKV(context.Background(), testKeyPrefix, KV{
		key:   testKey1,
		value: testValue1,
	}))
	assert.Nil(t, storeClient.SetWPrefixKV(context.Background(), testKeyPrefix, KV{
		key:   testKey2,
		value: testValue2,
	}))
}
