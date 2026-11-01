package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_tokenCmd = &cobra.Command{
	Use:   "token <command> [flags]",
	Short: "Manage runner controller tokens. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_tokenCmd).Standalone()

	runnerControllerCmd.AddCommand(runnerController_tokenCmd)
}
