package etcdwrap

import (
	"context"
	"testing"

	"github.com/TudorHulban/loginfo"
	"github.com/stretchr/testify/assert"
)

func TestCMD(t *testing.T) {
	shellCmd("ls", ".")
}

func TestSetSimple(t *testing.T) {
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
		key:   testKey1,
		value: testValue1,
	}))

	// test read key
	val1, errGet1 := storeClient.GetVByK(context.Background(), testKey1)
	assert.Nil(t, errGet1)
	assert.Equal(t, testValue1, val1)

	assert.Nil(t, storeClient.DeleteKVFrom(context.Background(), testKey1))

	// test read key
	_, errGet2 := storeClient.GetVByK(context.Background(), testKey1)
	assert.Error(t, errGet2)
}

// Target of test:
// a. insert
// b. update
// c. read
// d. read after reconnect
// e. close
func TestReconn(t *testing.T) {
	l1, errLog1 := loginfo.New(2)
	assert.Nil(t, errLog1)

	storeClient, err := NewETCDClient([]string{etcdURL}, l1)
	assert.Nil(t, err)
	assert.NotNil(t, storeClient)

	defer func() {
		assert.Nil(t, storeClient.TheStore.Close())
	}()

	// test set - a.
	assert.Nil(t, storeClient.SetKV(context.Background(), KV{
		key:   testKey1,
		value: testValue1,
	}))

	// test read key - c.
	val1, errGet1 := storeClient.GetVByK(context.Background(), testKey1)
	assert.Nil(t, errGet1)
	assert.Equal(t, testValue1, val1)

	// test update - b.
	assert.Nil(t, storeClient.SetKV(context.Background(), KV{
		key:   testKey1,
		value: testValue2,
	}))

	// test read key
	val2, errGet2 := storeClient.GetVByK(context.Background(), testKey1)
	assert.Nil(t, errGet2)
	assert.Equal(t, testValue2, val2)

	// close connection
	assert.Nil(t, storeClient.Close()) // e.

	l2, errLog2 := loginfo.New(2)
	assert.Nil(t, errLog2)

	// reconnect in order to recheck key
	reconClient, errRecon := NewETCDClient([]string{etcdURL}, l2)
	assert.Nil(t, errRecon)
	assert.NotNil(t, reconClient)
	defer reconClient.Close()

	// test read key - d.
	val3, errGet3 := reconClient.GetVByK(context.Background(), testKey1)
	assert.Nil(t, errGet3)
	assert.Equal(t, testValue2, val3)

	// delete key
	assert.Nil(t, reconClient.DeleteKVByK(context.Background(), testKey1))
}
