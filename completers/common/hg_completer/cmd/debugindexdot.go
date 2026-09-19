package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debugindexdotCmd = &cobra.Command{
	Use:    "debugindexdot",
	Short:  "-c|-m|FILE",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugindexdotCmd).Standalone()

	debugindexdotCmd.Flags().BoolP("changelog", "c", false, "open changelog")
	debugindexdotCmd.Flags().Bool("common", false, "also show common ancestors")
	debugindexdotCmd.Flags().String("dir", "", "open directory manifest")
	debugindexdotCmd.Flags().StringArrayP("include", "i", nil, "arbitrary graphviz statement to include")
	debugindexdotCmd.Flags().BoolP("manifest", "m", false, "open manifest")
	debugindexdotCmd.Flags().StringArrayP("rev", "r", nil, "only show these revisions")
	debugindexdotCmd.Flags().StringP("template", "T", "", "label template")
	rootCmd.AddCommand(debugindexdotCmd)

	carapace.Gen(debugindexdotCmd).FlagCompletion(carapace.ActionMap{
		"common":  action.ActionRevisions(),
		"include": carapace.ActionFiles(),
		"rev":     action.ActionRevisions(),
	})
}
