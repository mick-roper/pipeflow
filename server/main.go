package main

import (
	"github.com/mick-roper/pipeflow/server/config"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("server")

func main() {
	var loader config.Loader = config.NewEnvLoader()
	config, err := loader.Load()

	if err != nil {
		log.Panic(err)
	}

	log.Infof("server starting on port %v...", config.Port)
}
