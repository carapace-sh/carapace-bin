package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:     "pr",
	Short:   "Commands for creating and managing reviews on a forge, e.g. GitHub PRs or GitLab MRs",
	Aliases: []string{"review", "mr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prCmd).Standalone()

	prCmd.Flags().BoolP("draft", "d", false, "Whether to create reviews as a draft")
	prCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(prCmd)
}
