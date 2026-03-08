package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupListCmd = &cobra.Command{
	Use:   "list [group-spec]...",
	Short: "list groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupListCmd).Standalone()

	groupListCmd.Flags().Bool("available", false, "show only available groups")
	groupListCmd.Flags().String("contains-pkgs", "", "show only groups containing packages")
	groupListCmd.Flags().Bool("hidden", false, "show also hidden groups")
	groupListCmd.Flags().Bool("installed", false, "show only installed groups")

	groupCmd.AddCommand(groupListCmd)
}
