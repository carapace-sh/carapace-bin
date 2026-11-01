package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hash a local file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_hashCmd).Standalone()

	help_fileCmd.AddCommand(help_file_hashCmd)
}
