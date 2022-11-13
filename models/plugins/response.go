package plugins

type Response struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name string `json:"name"`
		Url  string `json:"browser_download_url"`
	} `json:"assets"`
}
