package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_verify_stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Verify repository state consistency (default)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_verify_stateCmd).Standalone()

	repository_help_verifyCmd.AddCommand(repository_help_verify_stateCmd)
}
