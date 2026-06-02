package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var home_buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a home-manager configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(home_buildCmd).Standalone()

	addHomeRebuildFlags(home_buildCmd)

	homeCmd.AddCommand(home_buildCmd)
}
