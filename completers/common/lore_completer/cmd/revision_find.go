package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_findCmd).Standalone()

	revision_findCmd.Flags().BoolP("help", "h", false, "Print help")
	revisionCmd.AddCommand(revision_findCmd)
}
