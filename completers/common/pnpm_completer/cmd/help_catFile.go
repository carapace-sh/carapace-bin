package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Prints the contents of a file based on the hash value stored in the index file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_catFileCmd).Standalone()

	helpCmd.AddCommand(help_catFileCmd)
}
