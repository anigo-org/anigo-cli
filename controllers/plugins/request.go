package plugins

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/FlamesX-128/anigo/models/plugins"
)

func FilterValidPlugins(ss []plugins.Response) map[string]plugins.Response {
	var data = map[string]plugins.Response{}

	for _, plugin := range ss {
		params := strings.Split(plugin.Name, "-")
		name := params[0]

		if len(params) != 4 || params[1] != PluginModelVersion {

			continue
		}

		if _, ok := data[name]; !ok {
			data[name] = plugin

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
