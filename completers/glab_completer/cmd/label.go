package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var labelCmd = &cobra.Command{
	Use:   "label <command> [flags]",
	Short: "Manage labels on remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labelCmd).Standalone()

	labelCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	rootCmd.AddCommand(labelCmd)

	carapace.Gen(labelCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(labelCmd),
	})
}
