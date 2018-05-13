package plugins

import "github.com/Jacobious52/expose/pkg/storage"

type help string

func (h *help) Setup() error {
	*h = "hello, world!"
	return nil
}

func (h help) Expose(storage.ReadStore) (string, error) {
	return string(h), nil
}
