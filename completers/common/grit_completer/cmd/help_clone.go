package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Copy a remote repository into a new directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cloneCmd).Standalone()

	helpCmd.AddCommand(help_cloneCmd)
}
