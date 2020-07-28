package etcdwrap

import "context"

// DeleteKVByK deletes KV by key.
func (s ETCDStore) DeleteKVByK(ctx context.Context, theK string) error {
	_, errDelete := s.TheStore.Delete(ctx, theK)
	return errDelete
}
