package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createFileCmd = &cobra.Command{
	Use:   "createFile",
	Short: "Create a file at the given path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createFileCmd).Standalone()
	carapace.Gen(createFileCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
	rootCmd.AddCommand(createFileCmd)
}
