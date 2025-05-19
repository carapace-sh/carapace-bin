package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:   "secret <command>",
	Short: "Manage GitHub secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretCmd).Standalone()

	secretCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(secretCmd)

	carapace.Gen(secretCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
