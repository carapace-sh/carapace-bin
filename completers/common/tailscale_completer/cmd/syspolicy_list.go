package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var syspolicy_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print effective policy settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syspolicy_listCmd).Standalone()

	syspolicyCmd.AddCommand(syspolicy_listCmd)
}
