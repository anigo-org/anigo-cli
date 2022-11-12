package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"plugin"
	"strconv"
	"strings"

	commandline "github.com/FlamesX-128/anigo/controllers/command-line"
	"github.com/FlamesX-128/anigo/controllers/surveys"
	"github.com/FlamesX-128/anigo/models"
	"golang.org/x/exp/maps"
)

var providers = map[string]models.Package{}

func searchOption(p models.Package) {
	var (
		entries = make(map[string]string)
		answer  string
		err     error
	)

	// Ask for novel to search.
	if answer, err = commandline.AskSearch(); err != nil {
		fmt.Println(err)

		return
	}

	// Search the novel.
	for _, novel := range p.Search(answer) {
		entries[novel.Title] = novel.Id
	}

	// Select the novel.
	if answer, err = commandline.AskNovel(maps.Keys(entries)); err != nil {
		fmt.Println(err)

		return
	}

	// Get novel info.
	info := p.Info(entries[answer])

	// Select episode.
	if answer, err = commandline.AskEpisode(len(info.Episodes)); err != nil {
		fmt.Println(err)

		return
	}

	epi, _ := strconv.Atoi(answer)

	// Get episode url, including quality.
	sources := p.Watch(info.Episodes[epi-1].Id)
	maps.Clear(entries)

	for _, source := range sources {
		entries[source.Quality] = source.Url
	}

	// Select quality.
	if answer, err = commandline.AskQuality(maps.Keys(entries)); err != nil {
		fmt.Println(err)

		return
	}

	// Play the episode.
	if err = commandline.Play(entries[answer]); err != nil {
		fmt.Println(err)

		return
	}
}

var menuOptions = map[string]func(string){
	"search": func(svc string) {
		searchOption(providers[svc])
	},
	"exit": func(_ string) {
		os.Exit(0)
	},
}

func LoadPlugin(path string) (models.Package, error) {
	plugin, err := plugin.Open(path)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	symbol, err := plugin.Lookup("Provider")

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	provider, ok := symbol.(models.Package)

	if !ok {
		return nil, fmt.Errorf("invalid package")
	}

	return provider, nil
}

//

func init() {
	filepath.Walk("plugins", func(s string, info os.FileInfo, err error) error {
		dir, _ := os.Getwd()

		if err == nil && strings.HasSuffix(info.Name(), ".so") {
			ss := path.Join(dir, s)

			log.Println("Loading", ss)

			if provider, err := LoadPlugin(ss); err == nil {
				providers[provider.Name()] = provider
			} else {
				log.Println("Failed to load", ss)
			}
		}

		return nil
	})
}

func main() {
	for {
		ans, err := surveys.Select("Select a provider", maps.Keys(providers))

		if err != nil {
			log.Panicln(err)
		}

		answer, err := surveys.Select("Select an option", maps.Keys(menuOptions))

		if err != nil {
			log.Panicln(err)
		}

		menuOptions[answer](ans)
	}
}
