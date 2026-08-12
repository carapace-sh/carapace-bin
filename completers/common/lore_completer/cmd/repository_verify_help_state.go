package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_verify_help_stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Verify repository state consistency (default)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_verify_help_stateCmd).Standalone()

	repository_verify_helpCmd.AddCommand(repository_verify_help_stateCmd)
}
