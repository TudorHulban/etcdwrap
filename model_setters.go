package etcdwrap

import (
	"context"
	"errors"

	"github.com/TudorHulban/loginfo"
	"go.etcd.io/etcd/clientv3"
)

// Set sets or updates key in store.
func (s ETCDStore) Set(ctx context.Context, theKV KV) error {
	_, errSet := s.TheStore.Put(ctx, theKV.key, theKV.value)
	return errSet
}
