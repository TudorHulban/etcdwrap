package etcdwrap

import (
	"context"
)

// GetVByK fetches value from store based on passed key.
func (s ETCDStore) GetVByK(ctx context.Context, theK string) (string, error) {
	result, errGet := s.TheStore.Get(ctx, theK)
	return string(result.Kvs[0].Value), errGet
}
