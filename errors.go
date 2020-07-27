package etcdwrap

import (
	"errors"
)

/*
File contains custom errors for etcd client and error handling for these errors.
These errors should be used with error wrapping.
*/

var (
	ErrorEtcdContextCancelled        = errors.New("ctx cancelled by another routine")
	ErrorEtcdContextDeadlineExceeded = errors.New("ctx deadline exceeded")
	ErrorEtcdClientSide              = errors.New("client side error")
	ErrorEtcdClusterSide             = errors.New("cluster side error")
)
