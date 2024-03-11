package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var org_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List organizations for the authenticated user.",
	GroupID: "General commands",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(org_listCmd).Standalone()

	org_listCmd.Flags().StringP("limit", "L", "", "Maximum number of organizations to list")
	orgCmd.AddCommand(org_listCmd)
}
