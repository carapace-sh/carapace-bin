package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Check Teleport config on the working Git directory. Or provide an action ('update' or 'reset') to configure the Git repo.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_configCmd).Standalone()

	gitCmd.AddCommand(git_configCmd)
}
