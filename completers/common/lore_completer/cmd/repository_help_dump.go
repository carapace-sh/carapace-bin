package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump repository state information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_dumpCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_dumpCmd)
}
