package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_stackCmd = &cobra.Command{
	Use:   "stack REPO",
	Short: "Configure Git stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_stackCmd).Standalone()

	gitCmd.AddCommand(git_stackCmd)
}
