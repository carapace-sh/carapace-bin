package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var home_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Build and activate a home-manager configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(home_switchCmd).Standalone()

	addHomeRebuildFlags(home_switchCmd)

	homeCmd.AddCommand(home_switchCmd)
}
