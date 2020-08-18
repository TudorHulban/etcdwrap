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

## Operations
### Cluster management
Version 
```
./etcdctl version
curl -L http://127.0.0.1:2379/version  # for Endpoint

```

Members List
```
./etcdctl --write-out=table --endpoints=localhost:2379 member list
```

### Key management
```
./etcdctl put key1 value1  # retrieved w ./etcdctl get key1 or curl http://127.0.0.1:2379/v2/keys/key1
./etcdctl put key2 value2  # retrieved w ./etcdctl get key2
./etcdctl put key3 value3 
./etcdctl put key4 value4 

./etcdctl put k3 value3 

./etcdctl get --from-key key # returns key1, key2
./etcdctl get --prefix k     # returns all keys starting with k

./etcdctl del k3             # check w ./etcdctl get k3
```
### Monitoring
```
https://etcd.io/docs/v3.4.0/op-guide/monitoring/
```


## Methods Exposed
The original implementation of etcd client is exposed. Besides this:<br/>
a. insert string / string key value<br/>
b. get string / string key value<br/>
c. insert several string / string key value pairs<br/>
d. get string / string key value<br/>




## Resources
```
https://etcd.io/docs/v2/api/   # v2!!
https://github.com/etcd-io/etcd/tree/master/etcdctl
https://godoc.org/github.com/coreos/etcd/clientv3
https://www.compose.com/articles/utilizing-etcd3-with-go/
https://github.com/etcd-io/etcd/blob/master/Documentation/learning/data_model.md
```

