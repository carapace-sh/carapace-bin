package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var labelCmd = &cobra.Command{
	Use:   "label <command> [flags]",
	Short: "Manage labels on remote.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labelCmd).Standalone()

	labelCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(labelCmd)

	carapace.Gen(labelCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(labelCmd),
	})
}
