package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Run a full garbage collection pass on the local repository store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_gcCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_gcCmd)
}
