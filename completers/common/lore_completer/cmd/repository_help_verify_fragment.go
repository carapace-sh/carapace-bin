package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_verify_fragmentCmd = &cobra.Command{
	Use:   "fragment",
	Short: "Verify a specific fragment in the local store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_verify_fragmentCmd).Standalone()

	repository_help_verifyCmd.AddCommand(repository_help_verify_fragmentCmd)
}
