package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/supervisor"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Activates any updates in config for process/group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				supervisor.ActionGroups(rootCmd.Flag("configuration").Value.String()),
				supervisor.ActionProcesses(rootCmd.Flag("configuration").Value.String()),
			).ToA().FilterArgs()
		}),
	)
}
