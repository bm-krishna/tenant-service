package driver

import "context"

// Connection is contract for setuping the store
type Connection interface {
	Provider(ctx context.Context, store string, config map[string]string) error
}
