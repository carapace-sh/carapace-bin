package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugpathcopiesCmd = &cobra.Command{
	Use:    "debugpathcopies",
	Short:  "hg debugpathcopies REV1 REV2 [FILE]",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugpathcopiesCmd).Standalone()

	debugpathcopiesCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	debugpathcopiesCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	rootCmd.AddCommand(debugpathcopiesCmd)

	carapace.Gen(debugpathcopiesCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
	})
}
