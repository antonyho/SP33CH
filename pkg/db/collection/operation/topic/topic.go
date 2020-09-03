// Package topic contains interface for topic operations,
// which should be implemented on all 'topic' database accessor.
package topic

// Operations for 'topic' database collection accessor interface
type Operations interface {
	List() ([]string, error)
	Create(topic string) error
}
