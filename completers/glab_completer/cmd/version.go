package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show version information for glab.",
	Aliases: []string{"v"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	rootCmd.AddCommand(versionCmd)
}
