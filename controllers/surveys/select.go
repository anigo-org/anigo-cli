package surveys

import (
	"github.com/AlecAivazis/survey/v2"
)

func Select(message string, options []string) (string, error) {
	err := survey.AskOne(&survey.Select{
		Message: message,
		Options: options,
	}, &message)

	return message, err
}
