package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:                "exec",
	Short:              "Executes a shell command in scope of a project",
	GroupID:            "run",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(execCmd).Standalone()

	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(execCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			command := c.Args[0]
			c.Args = c.Args[1:]
			return bridge.ActionCarapaceBin(command).Invoke(c).ToA()
		}),
	)
}
