package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:     "stack <command> [flags]",
	Short:   "Create, manage, and work with stacked diffs. (EXPERIMENTAL.)",
	Aliases: []string{"stacks"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stackCmd).Standalone()

	stackCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(stackCmd)
}
