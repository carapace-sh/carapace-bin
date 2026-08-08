package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupListCmd = &cobra.Command{
	Use:   "list [options] [<group-spec>...]",
	Short: "list comps groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupListCmd).Standalone()

	groupListCmd.Flags().Bool("available", false, "List available groups")
	groupListCmd.Flags().String("contains-pkgs", "", "Filter by packages in group")
	groupListCmd.Flags().Bool("hidden", false, "Include hidden groups")
	groupListCmd.Flags().Bool("installed", false, "List installed groups")

	groupCmd.AddCommand(groupListCmd)
}
