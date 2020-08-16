package etcdwrap

import "errors"

/*
File contains custom errors for etcd client and error handling for these errors.
These errors should be used with error wrapping.
*/

var errorNoValues = errors.New("no values retruned by client")

const errorClientSide = "could not perform operation"
