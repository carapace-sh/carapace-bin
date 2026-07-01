package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the launchd version",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	rootCmd.AddCommand(versionCmd)
}
