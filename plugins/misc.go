package plugins

import (
	"github.com/FlamesX-128/anigo-plugins/models"
	"github.com/FlamesX-128/anigo/plugins/exit"
)

var Plugin = models.Plugin{
	Providers: map[string]models.Provider{
		"[None] Exit": exit.Provider,
	},
}
