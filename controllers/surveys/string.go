package surveys

import (
	"github.com/manifoldco/promptui"
)

func String(label string, validate func(string) error) (string, error) {
	prompt := promptui.Prompt{
		Validate: validate,
		Label:    label,
	}

	return prompt.Run()
}
