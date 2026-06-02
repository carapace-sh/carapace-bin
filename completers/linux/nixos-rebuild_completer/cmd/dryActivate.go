package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dryActivateCmd = &cobra.Command{
	Use:   "dry-activate",
	Short: "Build the new configuration, but show what changes would occur on activation instead of activating the configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dryActivateCmd).Standalone()
	rootCmd.AddCommand(dryActivateCmd)
}
