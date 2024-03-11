package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/supervisor"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Reload config and add/remove as necessary, and will restart affected programs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				supervisor.ActionGroups(rootCmd.Flag("configuration").Value.String()),
				carapace.ActionValues("all"),
			).ToA().FilterArgs()
		}),
	)
}
