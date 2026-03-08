package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check for problems in the package database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().Bool("dependencies", false, "show missing dependencies and conflicts")
	checkCmd.Flags().Bool("duplicates", false, "show duplicated packages")
	checkCmd.Flags().Bool("obsoleted", false, "show obsoleted packages")

	rootCmd.AddCommand(checkCmd)
}
