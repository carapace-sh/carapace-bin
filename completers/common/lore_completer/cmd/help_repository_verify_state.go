package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_verify_stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Verify repository state consistency (default)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_verify_stateCmd).Standalone()

	help_repository_verifyCmd.AddCommand(help_repository_verify_stateCmd)
}
