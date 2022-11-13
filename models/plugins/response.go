package plugins

type Response struct {
	Name  string `json:"tag_name"`
	Files []struct {
		Name string `json:"name"`
		Url  string `json:"browser_download_url"`
	} `json:"assets"`
}
