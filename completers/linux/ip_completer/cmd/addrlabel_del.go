package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addrlabel_delCmd = &cobra.Command{
	Use:     "del",
	Aliases: []string{"delete"},
	Short:   "delete an address label entry",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addrlabel_delCmd).Standalone()

	addrlabelCmd.AddCommand(addrlabel_delCmd)
}
