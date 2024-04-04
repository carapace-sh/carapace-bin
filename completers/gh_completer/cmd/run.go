package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run <command>",
	Short:   "View details about workflow runs",
	GroupID: "actions",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
