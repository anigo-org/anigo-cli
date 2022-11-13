package main

import (
	"log"
	"os"

	"github.com/FlamesX-128/anigo-plugins/models"
	"github.com/FlamesX-128/anigo/controllers/interfaces"
	"github.com/FlamesX-128/anigo/controllers/plugins"
	"github.com/FlamesX-128/anigo/controllers/surveys"
	"github.com/fatih/color"
	"golang.org/x/exp/maps"
)

var providers = map[string]models.Provider{}

func init() {
	var root string
	var err error

	if root, err = os.Getwd(); err != nil {
		log.Panicln(color.RedString("[Plugin] [Init] [Error] %s", err.Error()))

		return
	}

	// Load plugins.
	for _, plugin := range plugins.Init[models.Plugin](root, "Plugin") {
		// Register provider.
		for name, provider := range plugin.Providers {
			providers[name] = provider
		}

	}

}

func main() {
	for {
		answer, err := surveys.Select("Select an option", maps.Keys(interfaces.MenuOptions))

		if err != nil {
			log.Panicln(err)
		}

		if answer == "Exit" {
			interfaces.MenuOptions["exit"](nil, nil)
		}

		ans, err := surveys.Select("Select a provider", maps.Keys(providers))

		if err != nil {
			log.Panicln(err)
		}

		interfaces.MenuOptions[answer](providers[ans], nil)
	}
}
