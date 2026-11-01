package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_verify_fragmentCmd = &cobra.Command{
	Use:   "fragment",
	Short: "Verify a specific fragment in the local store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_verify_fragmentCmd).Standalone()

	help_repository_verifyCmd.AddCommand(help_repository_verify_fragmentCmd)
}
