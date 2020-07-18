package etcdwrap

import (
	"context"
	"testing"

	"github.com/TudorHulban/loginfo"
	"github.com/stretchr/testify/assert"
)

const etcdURL = "127.0.0.1:2379"
const testKey = "the key"
const testValue1 = "the value1"
const testValue2 = "the value2"

// tests insert and update of key
func TestSet(t *testing.T) {
	l, errLog := loginfo.New(2)
	assert.Nil(t, errLog)

	storeClient, err := NewETCDClient([]string{etcdURL}, l)
	assert.Nil(t, err)
	assert.NotNil(t, storeClient)

	defer func() {
		assert.Nil(t, storeClient.TheStore.Close())
	}()

	// test set
	assert.Nil(t, storeClient.SetKV(context.Background(), KV{
		key:   testKey,
		value: testValue1,
	}))

	// test read key
	val1, errGet1 := storeClient.GetVByK(context.Background(), testKey)
	assert.Nil(t, errGet1)
	assert.Equal(t, testValue1, val1)

	// test update
	assert.Nil(t, storeClient.SetKV(context.Background(), KV{
		key:   testKey,
		value: testValue2,
	}))

	// test read key
	val2, errGet2 := storeClient.GetVByK(context.Background(), testKey)
	assert.Nil(t, errGet2)
	assert.Equal(t, testValue2, val2)
}
