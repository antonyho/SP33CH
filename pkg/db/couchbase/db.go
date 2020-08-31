package db

import (
	"time"

	"github.com/couchbase/gocb/v2"
)

const (
	// BucketName on Couchbase database.
	BucketName = "sp33ch"
)

// Connect to database by connection string.
// It returns Couchbase cluster and any connection error any encountered.
func Connect(str, username, password string) (*gocb.Cluster, error) {
	return gocb.Connect(
		str, gocb.ClusterOptions{
			Username: username,
			Password: password,
		},
	)
}

// GetCollection with given name, bucket and scope.
// The scope can be used as defining the environment,
// for example: dev, test, prod.
// It returns the collection database object.
func GetCollection(name, bucket string, cluster *gocb.Cluster) (*gocb.Collection, error) {
	// TODO - Create scope when it does not exist
	// https://pkg.go.dev/github.com/couchbase/gocb/v2?tab=doc#CollectionManager.CreateScope
	// TODO - Create collection when it does not exist
	// https://pkg.go.dev/github.com/couchbase/gocb/v2?tab=doc#CollectionManager.CreateCollection
	b := cluster.Bucket(bucket)
	if err := b.WaitUntilReady(5*time.Second, nil); err != nil {
		return nil, err
	}
	return b.Collection(name), nil
}
