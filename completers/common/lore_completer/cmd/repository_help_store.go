package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Access the repository data store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_storeCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_storeCmd)
}
