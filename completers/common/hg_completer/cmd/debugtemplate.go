package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debugtemplateCmd = &cobra.Command{
	Use:    "debugtemplate",
	Short:  "parse and apply a template",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugtemplateCmd).Standalone()

	debugtemplateCmd.Flags().StringArrayP("define", "D", nil, "define template keyword")
	debugtemplateCmd.Flags().StringArrayP("rev", "r", nil, "apply template on changesets")
	rootCmd.AddCommand(debugtemplateCmd)

	carapace.Gen(debugtemplateCmd).FlagCompletion(carapace.ActionMap{
		"rev": action.ActionRevisions(),
	})
}
