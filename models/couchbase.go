package models

import (
	"github.com/couchbase/gocb"
	"os"
	"github.com/vitorleandroloureiro/ACL/adapter"
)


type DB interface {
	Get(key string) interface{}
}


func NewCouchbase() *adapter.Bucket {

	//REMOVE THIS
	os.Setenv("couchbaseConnection","couchbase://localhost");
	os.Setenv("bucketName","default");
	os.Setenv("bucketPassword","");

	connection := os.Getenv("couchbaseConnection")
	bucketName := os.Getenv("bucketName")
	bucketPassword := os.Getenv("bucketPassword")

	cluster, _ := gocb.Connect(connection)
	bucket, _ := cluster.OpenBucket(bucketName, bucketPassword)

	return &adapter.Bucket{
		CbBucket: bucket,
	}
}
