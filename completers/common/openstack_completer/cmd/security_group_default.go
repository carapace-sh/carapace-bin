package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_defaultCmd).Standalone()

	security_groupCmd.AddCommand(security_group_defaultCmd)
}
