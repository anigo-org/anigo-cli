package surveys

import "github.com/AlecAivazis/survey/v2"

func Confirm(message string, def bool) (bool, error) {
	var confirm bool

	err := survey.AskOne(&survey.Confirm{
		Message: message,
		Default: def,
	}, &confirm)

	return confirm, err
}
