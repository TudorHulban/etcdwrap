package etcdwrap

import (
	"context"
)

// SetKV sets or updates key in store.
func (s ETCDStore) SetKV(ctx context.Context, theKV KV) error {
	_, errSet := s.TheStore.Put(ctx, theKV.key, theKV.value)
	return errSet
}

// SetWPrefixKV sets or updates key in store.
func (s ETCDStore) SetWPrefixKV(ctx context.Context, theKeyPrefix string, theKV KV) error {
	_, errSet := s.TheStore.Put(ctx, theKeyPrefix+theKV.key, theKV.value)
	return errSet
}
