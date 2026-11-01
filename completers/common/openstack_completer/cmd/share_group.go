package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_groupCmd).Standalone()

	shareCmd.AddCommand(share_groupCmd)
}
