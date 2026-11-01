package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeFileCmd = &cobra.Command{
	Use:   "removeFile",
	Short: "Remove the file at the given path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeFileCmd).Standalone()
	carapace.Gen(removeFileCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
	rootCmd.AddCommand(removeFileCmd)
}
