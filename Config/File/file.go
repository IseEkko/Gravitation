package File

import "github.com/IseEkko/Gravitation/Config"

var _ Config.Opertion = (*file)(nil)

type file struct {
	path string
}

func (f file) Load() ([]*Config.Value, error) {
	return []*Config.Value{}, nil
}
