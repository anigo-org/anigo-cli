package plugins

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/fatih/color"
)

var GOOS = map[string]string{
	"windows": "dll",
	"darwin":  "dylib",
	"linux":   "so",
}

func Init[T interface{}](root string, folder string, symName string) (data []*T) {
	pluginFolder := filepath.Join(root, folder)
	//goos := GOOS[runtime.GOOS]

	fmt.Println(pluginFolder)

	filepath.Walk(pluginFolder, func(dir string, info fs.FileInfo, err error) error {
		//filePath := filepath.Join(dir, info.Name())

		if err != nil || info.IsDir() {

			return err
		}

		var value *T

		color.Cyan("[Plugin] [Loading]: %s", info.Name())

		if value, err = Load[T](dir, symName); err != nil {
			color.Red("[Plugin] [Error] [%s]: %s", info.Name(), err.Error())

			return err
		}

		color.Green("[Plugin] [Loaded]: %s\n\n", info.Name())
		data = append(data, value)

		return err
	})

	return
}
