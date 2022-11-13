package surveys

import (
	"github.com/AlecAivazis/survey/v2"
)

func Multi(message string, options []string) ([]string, error) {
	var selected []string

	err := survey.AskOne(&survey.MultiSelect{
		Message: message,
		Options: options,
	}, &selected)

	return selected, err
}
