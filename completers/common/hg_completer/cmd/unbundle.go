package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unbundleCmd = &cobra.Command{
	Use:     "unbundle",
	Short:   "apply one or more bundle files",
	GroupID: groups[group_change_import_export].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unbundleCmd).Standalone()

	unbundleCmd.Flags().BoolP("update", "u", false, "update to new branch head if changesets were unbundled")
	rootCmd.AddCommand(unbundleCmd)

	carapace.Gen(unbundleCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
