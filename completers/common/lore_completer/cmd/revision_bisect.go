package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_bisectCmd = &cobra.Command{
	Use:   "bisect",
	Short: "Binary search for a change introduced between start (exclusive) and end (inclusive.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_bisectCmd).Standalone()

	revision_bisectCmd.Flags().String("end", "", "The earliest revision known to have the change")
	revision_bisectCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_bisectCmd.Flags().String("start", "", "The latest revision known to not have the change")
	revision_bisectCmd.MarkFlagRequired("end")
	revision_bisectCmd.MarkFlagRequired("start")
	revisionCmd.AddCommand(revision_bisectCmd)
}
