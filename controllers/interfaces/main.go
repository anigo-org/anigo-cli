package interfaces

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/FlamesX-128/anigo-plugins/models"
	"github.com/FlamesX-128/anigo/controllers/interfaces/ask"
	"golang.org/x/exp/maps"
)

func Play(url string) error {
	return exec.Command("mpv", url).Run()
}

func Exit(_ models.Provider, _ interface{}) {
	os.Exit(0)
}

//

func GetNovelInfo(provider models.Provider, data string) models.Info {
	return provider.Info(data)
}

//

func Watch(provider models.Provider, data interface{}) {
	var novel = GetNovelInfo(provider, data.(string))

	var entries = make(map[string]string)
	var answer string
	var err error

	// Select episode.
	if answer, err = ask.EpisodeToSelect(len(novel.Episodes)); err != nil {

		return
	}

	episode, _ := strconv.Atoi(answer)

	// Get episode url, including quality.
	for _, source := range provider.Watch(novel.Episodes[episode-1].Id) {
		entries[source.Quality] = source.Url
	}

	// Select quality.
	if answer, err = ask.QualityToSelect(maps.Keys(entries)); err != nil {

		return
	}

	if err := Play(entries[answer]); err != nil {

		return
	}
}

var MenuSearchOptions = map[string]func(models.Provider, interface{}){
	"watch": Watch,
	"exit":  Exit,
}

func Search(provider models.Provider, _ interface{}) {
	var values = make(map[string]string)
	var answer2 string
	var answer string
	var err error

	// Ask for novel to search.
	if answer, err = ask.NovelToSearch(); err != nil {

		return
	}

	// Search the novel.
	for _, novel := range provider.Search(answer) {
		values[novel.Title] = novel.Id
	}

	// Select the novel.
	if answer, err = ask.NovelToSelect(maps.Keys(values)); err != nil {

		return
	}

	// Select menu option.
	if answer2, err = ask.ToSelect(maps.Keys(MenuSearchOptions)); err != nil {

		return
	}

	MenuSearchOptions[answer2](
		provider, values[answer],
	)
}

//

var MenuOptions = map[string]func(models.Provider, interface{}){
	"Search": Search,
	"Exit":   Exit,
}
