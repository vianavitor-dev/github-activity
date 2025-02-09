package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/vianavitor-dev/github-activity/models"
)

func GetEventsCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("", flag.ExitOnError),
		Execute: getEvent,
	}

	return cmd
}

var getEvent = func(c *Command, args []string) {

	if len(args) <= 0 {
		fmt.Print(fmt.Errorf("getEvent: there are no arguments to use"))
		os.Exit(0)
	}

	username := args[0]

	resp, err := http.Get("https://api.github.com/users/" + username + "/events")

	if err != nil {
		log.Fatalf("getEvent %q : %v", username, err)
	}

	defer resp.Body.Close()

	var events []models.Event

	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		log.Fatalf("getEvent %q : %v", username, err)
	}

	if len(events) == 0 {
		fmt.Print(fmt.Errorf("getEvent %q : no activity found", username))
		os.Exit(0)
	}

	var pushCount = 0

	for _, e := range events {
		var message string

		switch e.Type {
		case "PushEvent":
			pushCount = len(e.Pay.Commits)
			message = fmt.Sprintf("Pushed %d commits to %s\n", pushCount, e.Repo.Name)
		case "IssuesEvent":
			message = fmt.Sprintf("Opened a new issue in %s\n", e.Repo.Name)
		case "WatchEvent":
			message = fmt.Sprintf("Starred %s\n", e.Repo.Name)
		case "CreateEvent":
			message = fmt.Sprintf("New %s created in %s\n", e.Pay.RefType, e.Repo.Name)
		case "ForkEvent":
			message = fmt.Sprintf("Fork repository %s\n", e.Repo.Name)
		default:
			message = fmt.Sprintf("%s in %s", strings.Replace(e.Type, "Event", "", 1), e.Repo.Name)
		}

		fmt.Print(message)
	}
	os.Exit(0)
}
