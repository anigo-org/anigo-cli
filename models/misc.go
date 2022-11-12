package models

/*
type Package struct {
	Search func(string) gogoanime.Search
	Info   func(string) gogoanime.Info
	Watch  func(string) gogoanime.Watch
}*/

type Package interface {
	Search(string) []Anime
	Info(string) Info
	Watch(string) []Source
	Name() string
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
