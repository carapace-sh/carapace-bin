package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_fileCmd = &cobra.Command{
	Use:   "file",
	Short: "File commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_fileCmd).Standalone()

	helpCmd.AddCommand(help_fileCmd)
}
