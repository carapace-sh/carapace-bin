package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check [options]",
	Short: "check for problems in the packagedb",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().Bool("dependencies", false, "Show missing dependencies and conflicts")
	checkCmd.Flags().Bool("duplicates", false, "Show duplicated packages")
	checkCmd.Flags().Bool("obsoleted", false, "Show obsoleted packages")

	rootCmd.AddCommand(checkCmd)
}
