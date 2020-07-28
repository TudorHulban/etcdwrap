package etcdwrap

import (
	"context"

	"go.etcd.io/etcd/clientv3"
)

// DeleteKVByK deletes KV by key.
func (s ETCDStore) DeleteKVByK(ctx context.Context, theK string) error {
	resp, errDelete := s.TheStore.Delete(ctx, theK)
	s.theLogger.Debug("Delete KV by key Response: ", *resp)
	return errDelete
}

// DeleteKVByPrefix deletes KV by prefix.
func (s ETCDStore) DeleteKVByPrefix(ctx context.Context, thePrefix string) error {
	resp, errDelete := s.TheStore.Delete(ctx, thePrefix, clientv3.WithPrefix())
	s.theLogger.Debug("Delete KVs by prefix Response: ", *resp)
	return errDelete
}
