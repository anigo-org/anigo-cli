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

	"github.com/FlamesX-128/anigo-plugins/models"
	commandline "github.com/FlamesX-128/anigo/controllers/command-line"
	"github.com/FlamesX-128/anigo/controllers/surveys"
	"golang.org/x/exp/maps"
)

var providers = map[string]models.Provider{}

func searchOption(p models.Provider) {
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

func LoadPlugin(path string) (*models.Plugin, error) {
	var (
		data interface{}
		err  error
	)

	if data, err = plugin.Open(path); err != nil {

		return nil, err
	}

	if data, err = data.(*plugin.Plugin).Lookup("Plugin"); err != nil {

		return nil, err
	}

	if data, ok := data.(*models.Plugin); ok {

		return data, nil
	}

	return nil, fmt.Errorf("unknown plugin type")
}

//

func init() {
	filepath.Walk("anigo-plugins", func(dir string, info os.FileInfo, err error) error {
		dirPath, _ := os.Getwd()

		if err == nil && strings.HasSuffix(info.Name(), ".so") {
			filePath := path.Join(dirPath, dir)

			fmt.Printf("Loading plugin %s\n", filePath)

			if data, err := LoadPlugin(filePath); err == nil {
				for name, provider := range data.Providers {
					providers[name] = provider
				}

			} else {
				fmt.Printf("Error loading plugin %s: %s\n", filePath, err)

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
