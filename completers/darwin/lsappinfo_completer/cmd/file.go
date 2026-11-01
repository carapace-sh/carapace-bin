package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Open the file at path and read lines as commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fileCmd).Standalone()
	carapace.Gen(fileCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
	rootCmd.AddCommand(fileCmd)
}
