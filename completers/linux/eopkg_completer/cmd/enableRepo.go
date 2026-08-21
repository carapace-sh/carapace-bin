package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enableRepoCmd = &cobra.Command{
	Use:     "enable-repo <name>",
	Aliases: []string{"er"},
	Short:   "enable a previously disabled repository by name",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableRepoCmd).Standalone()

	rootCmd.AddCommand(enableRepoCmd)
}
