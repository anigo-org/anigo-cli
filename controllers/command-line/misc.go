package commandline

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/FlamesX-128/anigo/controllers/surveys"
)

func AskSearch() (string, error) {
	return surveys.String("Search for a novel", nil)
}

func AskNovel(ss []string) (string, error) {
	switch len(ss) {
	case 0:
		return "", fmt.Errorf("no novels found")
	case 1:
		return ss[0], nil
	}

	return surveys.Select("Select a novel", ss)
}

func AskEpisode(max int) (string, error) {
	switch max {
	case 0:
		return "", fmt.Errorf("no episodes found")
	case 1:
		return "1", nil
	}

	return surveys.String(
		fmt.Sprintf("Select an episode (1-%d)", max),
		func(s string) error {
			if i, err := strconv.Atoi(s); err != nil || i < 1 || i > max {
				return fmt.Errorf("invalid episode number")
			}

			return nil
		},
	)
}

func AskQuality(ss []string) (string, error) {
	switch len(ss) {
	case 0:
		return "", fmt.Errorf("no qualities found")
	case 1:
		return ss[0], nil
	}

	return surveys.Select("Select a quality", ss)
}

func Play(url string) error {
	return exec.Command("mpv", url).Run()
}
