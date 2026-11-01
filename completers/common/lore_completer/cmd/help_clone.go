package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a remote repository into the given path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cloneCmd).Standalone()

	helpCmd.AddCommand(help_cloneCmd)
}
