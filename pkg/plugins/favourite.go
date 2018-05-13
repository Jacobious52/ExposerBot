package plugins

import (
	"bytes"

	"github.com/Jacobious52/expose/pkg/storage"
)

type favorite struct{}

func (f *favorite) Setup() error {
	return nil
}

func (f favorite) Expose(s storage.ReadStore) (string, error) {
	var b bytes.Buffer
	users := s.GetUsers()

	for _, user := range users {
		if table, ok := s.GetTable(user); ok {
			if len(table) == 0 {
				continue
			}
			b.WriteString(user)
			b.WriteString(": ")
			b.WriteString(table.Ordered()[0].Word)
		}
	}

	return b.String(), nil
}
