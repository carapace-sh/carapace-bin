package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var debugfilesetCmd = &cobra.Command{
	Use:    "debugfileset",
	Short:  "parse and apply a fileset specification",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugfilesetCmd).Standalone()

	debugfilesetCmd.Flags().Bool("all-files", false, "test files from all revisions and working directory")
	debugfilesetCmd.Flags().StringP("rev", "r", "", "apply the filespec on this revision")
	debugfilesetCmd.Flags().BoolP("show-matcher", "s", false, "print internal representation of matcher")
	debugfilesetCmd.Flags().StringArrayP("show-stage", "p", nil, "print parsed tree at the given stage")
	rootCmd.AddCommand(debugfilesetCmd)

	carapace.Gen(debugfilesetCmd).FlagCompletion(carapace.ActionMap{
		"rev": hg.ActionRevisions(),
	})
}
