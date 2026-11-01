package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:     "alias",
	Aliases: []string{"aliases", "ln"},
	Short:   "Interact with deployment aliases",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliasCmd).Standalone()

	rootCmd.AddCommand(aliasCmd)
}
