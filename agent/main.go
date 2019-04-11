package main

import "github.com/op/go-logging"

var log = logging.MustGetLogger("agent")

func main() {
	log.Info("agent started!")
}
