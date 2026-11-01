package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var disableRepoCmd = &cobra.Command{
	Use:     "disable-repo <name>",
	Aliases: []string{"dr"},
	Short:   "disable a system repository",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableRepoCmd).Standalone()

	rootCmd.AddCommand(disableRepoCmd)
}
