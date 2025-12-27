package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unitFilesCmd = &cobra.Command{
	Use:   "unit-files",
	Short: "List files and symlinks for units",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unitFilesCmd).Standalone()

	rootCmd.AddCommand(unitFilesCmd)
}
