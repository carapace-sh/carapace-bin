package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var writePIDToFileCmd = &cobra.Command{
	Use:   "writePIDToFile",
	Short: "Write the current pid to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(writePIDToFileCmd).Standalone()
	carapace.Gen(writePIDToFileCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
	rootCmd.AddCommand(writePIDToFileCmd)
}
