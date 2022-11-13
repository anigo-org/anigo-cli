package surveys

import "github.com/AlecAivazis/survey/v2"

func String(message string, validate func(interface{}) error) (string, error) {
	err := survey.AskOne(&survey.Input{
		Message: message,
	}, &message, survey.WithValidator(validate))

	return message, err
}
