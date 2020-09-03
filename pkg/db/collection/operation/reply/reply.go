// Package reply contains interface for reply operations,
// which should be implemented on all 'reply' database accessor.
package reply

import (
	"github.com/antonyho/SP33CH/pkg/db/collection/model"
)

// Operations for 'reply' database collection accessor interface
type Operations interface {
	Get(post string) ([]model.Reply, error)
	Post(reply *model.Reply) error
}
