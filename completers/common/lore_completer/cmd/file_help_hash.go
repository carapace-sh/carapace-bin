package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash a local file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_hashCmd).Standalone()

	file_helpCmd.AddCommand(file_help_hashCmd)
}
