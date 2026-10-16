package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backends_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List built-in backends",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backends_lsCmd).Standalone()

	backendsCmd.AddCommand(backends_lsCmd)
}
