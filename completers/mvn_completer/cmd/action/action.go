package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/mvn"
	"github.com/spf13/cobra"
)

func ActionProjects(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return mvn.ActionProjects(cmd.Flag("file").Value.String())
	})
}
