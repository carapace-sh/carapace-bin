package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_obliterateCmd = &cobra.Command{
	Use:   "obliterate",
	Short: "Obliterate a file or fragment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_obliterateCmd).Standalone()

	file_helpCmd.AddCommand(file_help_obliterateCmd)
}
