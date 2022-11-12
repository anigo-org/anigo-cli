package surveys

import (
	"github.com/manifoldco/promptui"
)

func Select(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Items: items,
		Label: label,
	}

	_, resp, err := prompt.Run()

	return resp, err
}
