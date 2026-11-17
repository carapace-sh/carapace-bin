package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var homeUnsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Follow this machine again",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(homeUnsetCmd).Standalone()

	homeCmd.AddCommand(homeUnsetCmd)
}
