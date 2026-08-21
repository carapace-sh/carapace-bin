package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_containsCmd = &cobra.Command{
	Use:   "contains",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_containsCmd).Standalone()

	groupCmd.AddCommand(group_containsCmd)
}
