// Package topic provides database operations for Topic model
package topic

import (
	"strings"

	"github.com/antonyho/SP33CH/pkg/db/collection/model"
	db "github.com/antonyho/SP33CH/pkg/db/couchbase"
	gocb "github.com/couchbase/gocb/v2"
)

const (
	// CollectionName on database
	CollectionName = "topics"
)

// Collection accessor for 'topic'
type Collection struct {
	*gocb.Collection
}

// New 'topic' collection accessor instance pointer
func New(cluster *gocb.Cluster) (*Collection, error) {
	coll, err := db.GetCollection(CollectionName, db.BucketName, db.GetScope(), cluster)
	if err != nil {
		return nil, err
	}
	return &Collection{Collection: coll}, nil
}

// List all topic
func (c *Collection) List() ([]string, error) {
	result, err := c.Collection.
}

// Create a new topic
func (c *Collection) Create(topic string) error {
	_, err := c.Insert(strings.ToLower(topic), model.Topic{Name: topic}, nil)
	if err != nil {
		return err
	}
	return nil
}
