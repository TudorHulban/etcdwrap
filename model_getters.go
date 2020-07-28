package etcdwrap

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

// GetVByK fetches value from store based on passed key.
func (s ETCDStore) GetVByK(ctx context.Context, theK string) (string, error) {
	resp, errGet := s.TheStore.Get(ctx, theK)
	return string(resp.Kvs[0].Value), errGet
}

// GetKVByKPrefix Method fetches KVs per passed prefix.
func (s ETCDStore) GetKVByKPrefix(ctx context.Context, thePrefix string) ([]KV, error) {
	resp, errGet := s.TheStore.Get(ctx, thePrefix, clientv3.WithPrefix())
	s.theLogger.Debug("Get KV by prefix Response: ", *resp)

	result := make([]KV, len(resp.Kvs))
	for i, v := range resp.Kvs {
		result[i].key = string(v.Key)
		result[i].value = string(v.Value)
	}

	s.theLogger.Debug("GetKVByKPrefix Result: ", result)
	return result, errGet
}
