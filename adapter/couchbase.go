package adapter

import "github.com/couchbase/gocb"

type Bucket struct {
	CbBucket gocb.Bucket
}

func (b *Bucket) Get(key string) interface{} {
	value := new(interface{})
	b.CbBucket.Get(key, value)

	return *value
}
