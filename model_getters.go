package etcdwrap

/*
Response in form of:
{cluster_id:14841639068965178418 member_id:10276657743932975437 revision:14 raft_term:2  [] false}
*/

import (
	"context"

	"github.com/pkg/errors"
	"go.etcd.io/etcd/clientv3"
)

// GetVByK fetches value from store based on passed key.
func (s ETCDStore) GetVByK(ctx context.Context, theK string) (string, error) {
	resp, errGet := s.TheStore.Get(ctx, theK)
	if errGet != nil {
		return "", errors.WithMessage(errGet, errorClientSide)
	}
	s.theLogger.Debug("GetVByK Result: ", *resp)
	if resp.Kvs == nil {
		return "", errorNoValues
	}
	return string(resp.Kvs[0].Value), errGet
}

// GetKVByKPrefix Method fetches KVs per passed prefix.
func (s ETCDStore) GetKVByKPrefix(ctx context.Context, thePrefix string) ([]KV, error) {
	resp, errGet := s.TheStore.Get(ctx, thePrefix, clientv3.WithPrefix())
	if errGet != nil {
		return []KV{}, errors.WithMessage(errGet, errorClientSide)
	}
	if resp.Kvs == nil {
		return []KV{}, errorNoValues
	}

	result := make([]KV, len(resp.Kvs))
	for i, v := range resp.Kvs {
		result[i].key = string(v.Key)
		result[i].value = string(v.Value)
	}

	s.theLogger.Debug("GetKVByKPrefix Result: ", result)
	return result, nil
}

// GetKVByKRangeFrom Method fetches KVs per passed range.
func (s ETCDStore) GetKVByKRangeFrom(ctx context.Context, rangeFrom string) ([]KV, error) {
	resp, errGet := s.TheStore.Get(ctx, rangeFrom, clientv3.WithFromKey())
	if errGet != nil {
		return []KV{}, errors.WithMessage(errGet, errorClientSide)
	}
	if resp.Kvs == nil {
		return []KV{}, errorNoValues
	}

	result := make([]KV, len(resp.Kvs))
	for i, v := range resp.Kvs {
		result[i].key = string(v.Key)
		result[i].value = string(v.Value)
	}

	s.theLogger.Debug("GetKVByKRange Result: ", result)
	return result, errGet
}
