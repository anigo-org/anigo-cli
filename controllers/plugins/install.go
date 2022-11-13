package plugins

import (
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func Install(path string, url string) {
	resp, err := http.Get(url)

	if err != nil {
		color.Red("[Plugin] [Install] [Error]: %s", err.Error())

		return
	}

	defer resp.Body.Close()

	file, err := os.Create(path)

	if err != nil {
		color.Red("[Plugin] [Create] [Error]: %s", err.Error())

		return
	}

	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		color.Red("[Plugin] [Copy] [Error]: %s", err.Error())

		return
	}

	color.Green("[Plugin] [Installed]: %s", path)

}
