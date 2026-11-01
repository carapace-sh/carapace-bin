package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_verify_help_fragmentCmd = &cobra.Command{
	Use:   "fragment",
	Short: "Verify a specific fragment in the local store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_verify_help_fragmentCmd).Standalone()

	repository_verify_helpCmd.AddCommand(repository_verify_help_fragmentCmd)
}
