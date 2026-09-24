package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugwalkCmd = &cobra.Command{
	Use:    "debugwalk",
	Short:  "show how files match on given patterns",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugwalkCmd).Standalone()

	debugwalkCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	debugwalkCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	rootCmd.AddCommand(debugwalkCmd)

	carapace.Gen(debugwalkCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
	})
}
