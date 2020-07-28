# etcdwrap
Implementing basic methods for interacting with etcd

## Etcd installation
```
https://github.com/TudorHulban/Deploy2020/blob/master/16_KVStores/01_etcd.md
```

## Client installation 
```
go get go.etcd.io/etcd/clientv3
```

## CLI Operations
```
./etcdctl put key1 value1  # retrieved w ./etcdctl get key1
./etcdctl put key2 value2  # retrieved w ./etcdctl get key2
./etcdctl put key3 value3 
./etcdctl put key4 value4 

./etcdctl put k3 value3 

./etcdctl get --from-key key # returns key1, key2
./etcdctl get --prefix k     # returns all keys starting with k

./etcdctl del k3             # check w ./etcdctl get k3
```

## Methods Exposed
The original implementation of etcd client is exposed. Besides this:<br/>
a. insert string / string key value<br/>
b. get string / string key value<br/>
c. insert several string / string key value pairs<br/>
d. get string / string key value<br/>




## Resources
```
https://github.com/etcd-io/etcd/tree/master/etcdctl
https://godoc.org/github.com/coreos/etcd/clientv3
https://www.compose.com/articles/utilizing-etcd3-with-go/
```

