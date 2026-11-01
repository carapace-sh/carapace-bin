package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addrlabel_flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "flush all address labels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addrlabel_flushCmd).Standalone()

	addrlabelCmd.AddCommand(addrlabel_flushCmd)
}
