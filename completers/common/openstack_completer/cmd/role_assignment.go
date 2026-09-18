package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var role_assignmentCmd = &cobra.Command{
	Use:   "assignment",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(role_assignmentCmd).Standalone()

	roleCmd.AddCommand(role_assignmentCmd)
}
