package main

import (
	"os"

	"github.com/vianavitor-dev/github-activity/cmd"
)

func main() {
	var c = cmd.GetEventsCommand()

	c.Init(os.Args[1:])
	c.Run()
}
