package etcdwrap

import (
	"context"
	"testing"

	"github.com/TudorHulban/loginfo"
	"github.com/stretchr/testify/assert"
)

const etcdURL = "127.0.0.1:2379"
const testKey = "the key"
const testValue = "the value"

func TestSet(t *testing.T) {
	l, errLog := loginfo.New(2)
	assert.Nil(t, errLog)

	storeClient, err := NewETCDClient([]string{etcdURL}, l)
	assert.Nil(t, err)
	assert.NotNil(t, storeClient)

	defer func() {
		assert.Nil(t, storeClient.TheStore.Close())
	}()

	// test insert
	assert.Nil(t, storeClient.SetKV(context.Background(), KV{
		key:   testKey,
		value: testValue,
	}))

	// test read key
	val, errGet := storeClient.GetVByK(context.Background(), testKey)
	assert.Nil(t, errGet)
	assert.Equal(t, testValue, val)
}
