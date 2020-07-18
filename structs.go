package etcdwrap

import (
	"github.com/TudorHulban/loginfo"
	"go.etcd.io/etcd/clientv3"
)

// ETCDStore Concentrates information defining a KV store.
type ETCDStore struct {
	theLogger loginfo.LogInfo // logger needed only for package logging
	TheStore  *clientv3.Client
}
