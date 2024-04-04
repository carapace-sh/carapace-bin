package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var labelCmd = &cobra.Command{
	Use:   "label <command>",
	Short: "Manage labels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labelCmd).Standalone()

	labelCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(labelCmd)

	carapace.Gen(labelCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
