package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_versionCmd).Standalone()

	helpCmd.AddCommand(help_versionCmd)
}
