package plugins

import (
	"path/filepath"

	"github.com/FlamesX-128/anigo/controllers/filesystem"
	"github.com/FlamesX-128/anigo/controllers/interfaces/ask"
	"github.com/FlamesX-128/anigo/models/plugins"
	"github.com/fatih/color"
	"golang.org/x/exp/maps"
)

func Init[T interface{}](root string, symName string) (data []*T) {
	var path string = filepath.Join(root, AnigoFolderName, PluginsFolderName)
	var err error

	// Check if the plugins folder exists.
	if ok, _ := filesystem.CheckIfPathExists(path); !ok {
		var data map[string]plugins.Response
		var answer bool

		if _, err = filesystem.CreateFolderIfNotExists(path); err != nil {
			color.Red("[Plugin] [Init] [Error] %s", err.Error())

			return nil
		}

		// Ask user if he wants to install plugins.
		if answer, err = ask.InstallSomePlugin(); err != nil {
			color.Red("[Plugin] [Selection] [Error] %s", err.Error())

			return nil
		}

		// Install plugins if user wants to.
		if answer {
			// Get plugins from github.
			if data, err = Request(); err != nil {
				color.Red("[Plugin] [Request] [Error] %s", err.Error())

				return nil
			}

			MultiInstall(path, maps.Values(data))
		}

	}

	// Load plugins.
	return Load[T](path, symName)
}
