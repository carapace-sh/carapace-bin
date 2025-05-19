package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Displays information related to a deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()

	rootCmd.AddCommand(inspectCmd)

	carapace.Gen(inspectCmd).PositionalCompletion(
		action.ActionDeployments(inspectCmd),
	)
}
