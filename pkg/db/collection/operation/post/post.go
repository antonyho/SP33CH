// Package post contains interface for post operations,
// which should be implemented on all 'post' database accessor.
package post

import (
	"github.com/antonyho/SP33CH/pkg/db/collection/model"
)

// Operations for 'post' database collection accessor interface
type Operations interface {
	List(topic string) ([]model.Post, error)
	Post(thread *model.Post) error
}
