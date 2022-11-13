package plugins

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/FlamesX-128/anigo/models/plugins"
)

func Search() (_ map[string]plugins.Response, err error) {
	var resp *http.Response
	var req *http.Request

	if req, err = http.NewRequest("GET", plugins.URL, nil); err != nil {

		return
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	if resp, err = http.DefaultClient.Do(req); err != nil {

		return
	}

	fmt.Println("1:")

	bodyBytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println("2:")

	var response []plugins.Response

	fmt.Println("3:")

	if err = json.Unmarshal(bodyBytes, &response); err != nil {

		return
	}

	var data = map[string]plugins.Response{}

	fmt.Println("4:", response)

	for _, plugin := range response {
		params := strings.Split(plugin.TagName, "-")

		fmt.Println(params, len(params))

		if len(params) == 4 && params[1] == PluginModelVersion {
			if _, ok := data[params[0]]; !ok {
				data[params[0]] = plugin
			}
		}
	}

	return data, nil
}
