package plugins

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/FlamesX-128/anigo/models/plugins"
	"github.com/hashicorp/go-version"
)

func FilterValidPlugins(ss []plugins.Response) map[string]plugins.Response {
	var data = map[string]plugins.Response{}

	for _, plugin := range ss {
		params := strings.Split(plugin.Name, "-")
		info := strings.Split(params[0], "@")
		skip := false

		if len(params) != 4 || params[1] != PluginModelVersion {

			continue
		}

		for name := range data {
			if !strings.HasPrefix(name, info[0]) {

				continue
			}

			v1, err := version.NewVersion(strings.Split(name, "@")[1])
			v2, err2 := version.NewVersion(info[1])

			if err != nil || err2 != nil {
				log.Panicln(err, err2)
			}

			if v1.GreaterThan(v2) {
				skip = true
			} else {
				delete(data, name)
			}

		}

		if !skip {
			data[params[0]] = plugin
		}

	}

	return data
}

func Request() (_ map[string]plugins.Response, err error) {
	var resp *http.Response
	var req *http.Request

	if req, err = http.NewRequest("GET", plugins.URL, nil); err != nil {

		return
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	if resp, err = http.DefaultClient.Do(req); err != nil {

		return
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response []plugins.Response

	if err = json.Unmarshal(bodyBytes, &response); err != nil {

		return
	}

	return FilterValidPlugins(response), nil
}
