package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enableRepoCmd = &cobra.Command{
	Use:     "enable-repo",
	Short:   "Enable a repository",
	Aliases: []string{"er"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableRepoCmd).Standalone()

	rootCmd.AddCommand(enableRepoCmd)
}
