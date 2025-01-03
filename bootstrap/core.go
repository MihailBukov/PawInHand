package bootstrap

import (
	"PawInHand/config"
	"go.uber.org/fx"
	"os"
)

var FXModule_Core = fx.Module(
	"core",
	fx.Provide(
		createApplicationConfiguration,
	),
)

func createApplicationConfiguration() (*config.ApplicationConfiguration, error) {
	return config.NewApplicationConfiguration(os.Getenv("ENVIRONMENT"))
}
