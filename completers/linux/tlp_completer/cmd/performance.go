package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var performanceCmd = &cobra.Command{
	Use:   "performance",
	Short: "Apply the performance profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(performanceCmd).Standalone()

	rootCmd.AddCommand(performanceCmd)

	carapace.Gen(performanceCmd).DashAnyCompletion(action.ActionConfigParams())
}
