package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff two revisions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_diffCmd).Standalone()

	revision_diffCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_diffCmd.Flags().StringSlice("path", nil, "Optional path in repository")
	revision_diffCmd.Flags().String("target", "", "Target revision to compare, by default the current revision")
	revision_diffCmd.Flags().String("targets", "", "Path to a targets file")
	revisionCmd.AddCommand(revision_diffCmd)
}
