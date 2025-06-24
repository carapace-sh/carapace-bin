package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unitPathsCmd = &cobra.Command{
	Use:   "unit-paths",
	Short: "List load directories for units",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unitPathsCmd).Standalone()

	rootCmd.AddCommand(unitPathsCmd)
}
