package exit

import (
	"os"

	"github.com/FlamesX-128/anigo-plugins/models"
)

type ProviderModel struct{}

func (p ProviderModel) Search(query string) []models.Anime {
	os.Exit(1)

	return nil
}

func (p ProviderModel) Info(id string) models.Info {
	os.Exit(1)

	return models.Info{}
}

func (p ProviderModel) Watch(id string) []models.Source {
	os.Exit(1)

	return nil
}

var Provider models.Provider = ProviderModel{}
