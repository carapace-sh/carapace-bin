package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Removes a deployment",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalCompletion(
		carapace.Batch(
			action.ActionProjects(rmCmd),
			action.ActionDeployments(rmCmd),
		).ToA(),
	)
}
