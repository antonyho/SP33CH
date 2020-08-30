package db

import (
	"time"

	"github.com/couchbase/gocb/v2"
)

const (
	// BucketName on Couchbase database
	BucketName = "sp33ch"
)

// Different Couchbase scopes in bucket
const (
	ScopeTest       = "test"
	ScopeProduction = "prod"
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

// Collection with given name, bucket and scope.
// The scope can be used as defining the environment,
// for example: dev, test, prod.
// It returns the collection database object.
func Collection(name, bucket, scope string, cluster *gocb.Cluster) (*gocb.Collection, error) {
	b := cluster.Bucket(bucket)
	if err := b.WaitUntilReady(5*time.Second, nil); err != nil {
		return nil, err
	}
	return b.Scope(scope).Collection(name), nil
}
