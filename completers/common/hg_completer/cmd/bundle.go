package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bundleCmd = &cobra.Command{
	Use:     "bundle",
	Short:   "create a bundle file",
	GroupID: groups[group_change_import_export].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundleCmd).Standalone()

	bundleCmd.Flags().BoolP("all", "a", false, "bundle all changesets in the repository")
	bundleCmd.Flags().String("base", "", "a base changeset assumed to be available at the destination")
	bundleCmd.Flags().StringArrayP("branch", "b", nil, "a specific branch you would like to bundle")
	bundleCmd.Flags().Bool("exact", false, "compute the base from the revision specified")
	bundleCmd.Flags().BoolP("force", "f", false, "run even when the destination is unrelated")
	bundleCmd.Flags().Bool("insecure", false, "do not verify server certificate (ignoring web.cacerts config)")
	bundleCmd.Flags().String("remotecmd", "", "specify hg command to run on the remote side")
	bundleCmd.Flags().StringArrayP("rev", "r", nil, "a changeset intended to be added to the destination")
	bundleCmd.Flags().StringP("ssh", "e", "", "specify ssh command to use")
	bundleCmd.Flags().StringP("type", "t", "", "bundle compression type to use")
	rootCmd.AddCommand(bundleCmd)

	carapace.Gen(bundleCmd).FlagCompletion(carapace.ActionMap{
		"base":      action.ActionRevisions(),
		"branch":    action.ActionBranches(),
		"remotecmd": carapace.ActionExecutables(),
		"rev":       action.ActionRevisions(),
		"ssh":       carapace.ActionExecutables(),
	})
}
