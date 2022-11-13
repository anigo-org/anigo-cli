package ask

import "github.com/FlamesX-128/anigo/controllers/surveys"

func InstallSomePlugin() (bool, error) {
	return surveys.Confirm(
		"No plugins found, do you want to install some?", false,
	)
}
