package plugins

import (
	"io/fs"
	"path/filepath"

	"github.com/fatih/color"
)

func Load[T interface{}](folder string, symName string) (data []*T) {
	filepath.Walk(folder, func(dir string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {

			return err
		}

		var value *T

		color.Cyan("[Plugin] [Loading]: %s", info.Name())

		if value, err = Open[T](dir, symName); err != nil {
			color.Red("[Plugin] [Error] [%s]: %s", info.Name(), err.Error())

			return err
		}

		color.Green("[Plugin] [Loaded]: %s\n\n", info.Name())
		data = append(data, value)

		return err
	})

	return
}
