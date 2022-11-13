package plugins

import (
	"errors"
	"plugin"
)

func Open[T interface{}](path string, symName string) (data *T, err error) {
	var entity *plugin.Plugin
	var symbol plugin.Symbol

	if entity, err = plugin.Open(path); err != nil {
		return nil, err
	}

	if symbol, err = entity.Lookup(symName); err != nil {
		return nil, err
	}

	if data, ok := symbol.(*T); ok {
		return data, nil
	}

	return nil, errors.New(
		"plugin symbol does not have the correct type",
	)
}
