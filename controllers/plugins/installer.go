package plugins

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/FlamesX-128/anigo/models/plugins"
	"github.com/fatih/color"
)

func Installer(path string, url string) (err error) {
	var resp *http.Response
	var data []byte

	if resp, err = http.Get(url); err != nil {
		return
	}

	defer resp.Body.Close()

	if data, err = io.ReadAll(resp.Body); err != nil {
		return
	}

	if err = os.WriteFile(path, data, 0644); err != nil {
		return
	}

	return nil
}

func Install(path string, plugin plugins.Response) {
	for _, file := range plugin.Files {
		color.Cyan("[Plugin] [Installing] %s", file.Name)

		if !strings.Contains(file.Name, (runtime.GOOS + "-" + runtime.GOARCH)) {

			continue
		}

		if err := Installer(filepath.Join(path, file.Name), file.Url); err != nil {
			color.Red("[Plugin] [Error]: %s", err.Error())

			continue
		}

		color.Green("[Plugin] [Installed] %s\n\n", file.Name)
	}
}

func MultiInstall(path string, plugins []plugins.Response) {
	for _, plugin := range plugins {
		Install(path, plugin)
	}
}
