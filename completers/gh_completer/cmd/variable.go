package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var variableCmd = &cobra.Command{
	Use:   "variable <command>",
	Short: "Manage GitHub Actions variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variableCmd).Standalone()

	variableCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(variableCmd)

	carapace.Gen(variableCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
