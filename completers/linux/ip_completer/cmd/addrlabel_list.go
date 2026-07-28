package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addrlabel_listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"show"},
	Short:   "list address label entries",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addrlabel_listCmd).Standalone()

	addrlabelCmd.AddCommand(addrlabel_listCmd)
}
