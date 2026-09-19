package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var manifestCmd = &cobra.Command{
	Use:     "manifest",
	Short:   "output the current or given revision of the project manifest",
	GroupID: groups[group_repository_maintenance].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifestCmd).Standalone()

	manifestCmd.Flags().Bool("all", false, "list files from all revisions")
	manifestCmd.Flags().StringP("rev", "r", "", "revision to display")
	manifestCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(manifestCmd)

	carapace.Gen(manifestCmd).FlagCompletion(carapace.ActionMap{
		"rev": action.ActionRevisions(),
	})
}
