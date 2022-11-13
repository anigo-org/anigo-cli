package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/FlamesX-128/anigo-plugins/models"
	commandline "github.com/FlamesX-128/anigo/controllers/command-line"
	"github.com/FlamesX-128/anigo/controllers/plugins"
	"github.com/FlamesX-128/anigo/controllers/surveys"
	"github.com/fatih/color"
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

func init() {
	root, err := os.Getwd()

	if err != nil {
		log.Panicln(err)
	}

	anigoFolder := filepath.Join(root, "anigo")

	anigoPluginFolder := filepath.Join(anigoFolder, "plugins")
	//anigoConfig := filepath.Join(anigoFolder, "config.json")

	if _, err := os.Stat(anigoFolder); os.IsNotExist(err) {
		/*if err := os.WriteFile(anigoConfig, []byte{}, 0644); err != nil {
			log.Panicln(err)
		}*/

		if err := os.MkdirAll(anigoPluginFolder, 0755); err != nil {
			log.Panicln(err)
		}

		data, err := plugins.Search()

		if err != nil {
			color.Red("[Plugin] [Search] [Error] %s", err)

			return
		}

		toInstall, err := commandline.AskPlugin(maps.Keys(data))

		if err != nil {
			color.Red("[Plugin] [Selection] [Error] %s", err)

			return
		}

		for _, plugin := range toInstall {
			for _, asset := range data[plugin].Assets {
				fmt.Println(asset.Name, (runtime.GOOS + "-" + runtime.GOARCH))
				if !strings.Contains(asset.Name, (runtime.GOOS + "-" + runtime.GOARCH)) {

					continue
				}

				plugins.Install(filepath.Join(anigoPluginFolder, plugin), asset.Url)
			}
		}

	}

	for _, plugin := range plugins.Init[models.Plugin](anigoFolder, "plugins", "Plugin") {
		for key, value := range plugin.Providers {
			providers[key] = value
		}

	}
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
