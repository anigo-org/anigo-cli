package ask

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/FlamesX-128/anigo/controllers/surveys"
)

func NovelToSearch() (string, error) {
	return surveys.String("what novel do you want to look for?", func(i interface{}) error {

		return nil
	})
}

func NovelToSelect(options []string) (string, error) {
	return surveys.Select("which novel do you want to select?", options)
}

func EpisodeToSelect(size int) (string, error) {
	switch size {
	case 0:
		return "1", errors.New("no episodes found")
	case 1:
		return "2", nil
	}

	message := fmt.Sprintf("which episode do you want to select? (1-%d)", size)

	return surveys.String(message, func(i interface{}) error {
		u, err := strconv.Atoi(i.(string))

		if err != nil || u < 1 || u > size {
			return errors.New("invalid episode number")
		}

		return nil
	})
}

func QualityToSelect(options []string) (string, error) {
	return surveys.Select("which quality do you want to select?", options)
}

func ToSelect(options []string) (string, error) {
	return surveys.Select("which do you want to select?", options)
}
