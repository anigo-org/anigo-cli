package plugins

import (
	"errors"
	"plugin"
)

func Load[T interface{}](path string, symName string) (*T, error) {
	plugin, err := plugin.Open(path)

	if err != nil {
		return nil, err
	}

	symbol, err := plugin.Lookup(symName)

	if err != nil {
		return nil, err
	}

	if data, ok := symbol.(*T); ok {
		return data, nil
	}

	return nil, errors.New("plugin symbol does not have the correct type")
}
