package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var phaseCmd = &cobra.Command{
	Use:     "phase",
	Short:   "set or show the current phase name",
	GroupID: groups[group_change_organization].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(phaseCmd).Standalone()

	phaseCmd.Flags().BoolP("draft", "d", false, "set changeset phase to draft")
	phaseCmd.Flags().BoolP("force", "f", false, "allow to move boundary backward")
	phaseCmd.Flags().BoolP("public", "p", false, "set changeset phase to public")
	phaseCmd.Flags().StringArrayP("rev", "r", nil, "target revision")
	phaseCmd.Flags().BoolP("secret", "s", false, "set changeset phase to secret")
	rootCmd.AddCommand(phaseCmd)

	carapace.Gen(phaseCmd).FlagCompletion(carapace.ActionMap{
		"rev": action.ActionRevisions(),
	})
}
