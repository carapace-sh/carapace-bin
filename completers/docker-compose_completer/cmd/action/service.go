package action

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type config struct {
	Services map[string]struct {
		Image string
		Build struct {
			Context    string
			Dockerfile string
		}
	}
	Volumes map[string]struct {
		Name string
	}
}

func actionConfig(f func(c config) carapace.Action) carapace.Action {
	// TODO support `--files` flag
	return carapace.ActionExecCommand("docker-compose", "convert", "--format", "json")(func(output []byte) carapace.Action {
		var c config
		if err := json.Unmarshal(output, &c); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(c)
	})
}

func ActionServices(cmd *cobra.Command) carapace.Action {
	return actionConfig(func(c config) carapace.Action {
		vals := make([]string, 0)
		for name, service := range c.Services {
			description := service.Image
			if strings.TrimSpace(description) == "" {
				description = fmt.Sprintf("%v/%v", service.Build.Context, service.Build.Dockerfile)
			}
			vals = append(vals, name, description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionContainers(cmd *cobra.Command, status string) carapace.Action {
	return carapace.ActionExecCommand("docker-compose", "ps", "--status", status)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		if lines[0] != "" {
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}
		return carapace.ActionValues()
	})
}
