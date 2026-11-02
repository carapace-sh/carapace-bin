package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var autoupdateCmd = &cobra.Command{
	Use:   "autoupdate",
	Short: "Auto-update pre-commit config to the latest repos' versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoupdateCmd).Standalone()

	autoupdateCmd.Flags().Bool("bleeding-edge", false, "Update to the bleeding edge of `HEAD` instead of the latest tagged version (the default behavior).")
	autoupdateCmd.Flags().String("color", "", "Whether to use color in output.")
	autoupdateCmd.Flags().StringP("config", "c", "", "Path to alternate config file")
	autoupdateCmd.Flags().Bool("freeze", false, "Store \"frozen\" hashes in `rev` instead of tag names")
	autoupdateCmd.Flags().StringP("jobs", "j", "", "Number of threads to use.")
	autoupdateCmd.Flags().String("repo", "", "Only update this repository -- may be specified multiple times.")

	carapace.Gen(autoupdateCmd).FlagCompletion(carapace.ActionMap{
		"color":  action.ActionColors(),
		"config": carapace.ActionFiles(".yaml", ".yml"),
	})

	rootCmd.AddCommand(autoupdateCmd)
}
