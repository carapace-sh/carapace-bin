package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker/compose"
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

func actionConfig(cmd *cobra.Command, f func(c config) carapace.Action) carapace.Action {
	return actionExecCompose(cmd, "convert", "--format", "json")(func(output []byte) carapace.Action {
		var c config
		if err := json.Unmarshal(output, &c); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(c)
	})
}

func ActionServices(cmd *cobra.Command) carapace.Action {
	files, err := cmd.Root().Flags().GetStringArray("file")
	if err != nil {
		return carapace.ActionMessage(err.Error())
	}
	return compose.ActionServices(files...)
}
