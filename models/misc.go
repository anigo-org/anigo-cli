package models

type Provider interface {
	Search(string) []Anime
	Info(string) Info
	Watch(string) []Source
}

type Plugin struct {
	Providers map[string]Provider
}

//

type Source struct {
	Quality string
	Url     string
}

type Episode struct {
	Number uint32
	Id     string
}

type Info struct {
	Episodes []Episode
	Image    string
}

type Anime struct {
	Title string
	Id    string
}
