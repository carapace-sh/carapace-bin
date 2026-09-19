package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var debugpickmergetoolCmd = &cobra.Command{
	Use:    "debugpickmergetool",
	Short:  "examine which merge tool is chosen for specified file",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugpickmergetoolCmd).Standalone()

	debugpickmergetoolCmd.Flags().Bool("changedelete", false, "emulate merging change and delete")
	debugpickmergetoolCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	debugpickmergetoolCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	debugpickmergetoolCmd.Flags().StringP("rev", "r", "", "check for files in this revision")
	debugpickmergetoolCmd.Flags().StringP("tool", "t", "", "specify merge tool")
	rootCmd.AddCommand(debugpickmergetoolCmd)

	carapace.Gen(debugpickmergetoolCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"rev":     hg.ActionRevisions(),
		"tool":    hg.ActionMergeTools(),
	})
}
