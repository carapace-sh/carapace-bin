package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload conda packages to various channels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_uploadCmd).Standalone()

	helpCmd.AddCommand(help_uploadCmd)
}
