package etcdwrap

import (
	"context"
)

// SetKV sets or updates key in store.
func (s ETCDStore) SetKV(ctx context.Context, theKV KV) error {
	_, errSet := s.TheStore.Put(ctx, theKV.key, theKV.value)
	return errSet
}
