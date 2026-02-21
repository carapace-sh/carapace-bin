package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var disableRepoCmd = &cobra.Command{
	Use:     "disable-repo",
	Short:   "Disable a repository",
	Aliases: []string{"dr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableRepoCmd).Standalone()

	rootCmd.AddCommand(disableRepoCmd)
}
