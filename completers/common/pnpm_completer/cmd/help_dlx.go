package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_dlxCmd = &cobra.Command{
	Use:   "dlx",
	Short: "Run a package in a temporary environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_dlxCmd).Standalone()

	helpCmd.AddCommand(help_dlxCmd)
}
