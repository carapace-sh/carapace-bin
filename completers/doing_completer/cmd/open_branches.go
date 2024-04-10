package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_branchesCmd = &cobra.Command{
	Use:   "branches",
	Short: "Open an overview of the repositories' branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_branchesCmd).Standalone()

	open_branchesCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_branchesCmd)
}
