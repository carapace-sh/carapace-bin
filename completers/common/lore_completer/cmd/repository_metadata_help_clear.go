package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_metadata_help_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata from the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_metadata_help_clearCmd).Standalone()

	repository_metadata_helpCmd.AddCommand(repository_metadata_help_clearCmd)
}
