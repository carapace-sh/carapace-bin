package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var syspolicy_reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Force a reload of policy settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syspolicy_reloadCmd).Standalone()

	syspolicyCmd.AddCommand(syspolicy_reloadCmd)
}
