package main

import (
	"os"

	"github.com/aethiopicuschan/argus"
)

func main() {
	logger := argus.NewLogger(os.Stdout, argus.WithColor())
	logger.Debug().Add("key", "value").Print()
	logger.Info().Add("key", "value").Print()
	logger.Warn().Add("key", "value").Print()
	logger.Error().Add("key", "value").Print()
}
