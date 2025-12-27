package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compareVersionsCmd = &cobra.Command{
	Use:   "compare-versions",
	Short: "VERSION2 Compare two version strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compareVersionsCmd).Standalone()

	rootCmd.AddCommand(compareVersionsCmd)
}
