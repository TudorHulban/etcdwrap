package etcdwrap

import (
	"errors"

	"github.com/TudorHulban/loginfo"
	"go.etcd.io/etcd/clientv3"
)

// NewETCDClient provides client for no security connections
func NewETCDClient(endpoints []string, extLogger loginfo.LogInfo) (ETCDStore, error) {
	nilStore := ETCDStore{
		theLogger: extLogger,
		TheStore:  nil,
	}

	if len(endpoints) == 0 {
		return nilStore, errors.New("empty endpoints list provided")
	}

	etcdClient, errClient := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if errClient != nil {
		return nilStore, errClient
	}
	defer etcdClient.Close()

	return ETCDStore{
		theLogger: extLogger,
		TheStore:  etcdClient,
	}, nil
}

// Close closes the store.
func (s ETCDStore) Close() error {
	return s.TheStore.Close()
}
