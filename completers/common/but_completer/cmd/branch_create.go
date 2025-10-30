package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new virtual branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_createCmd).Standalone()

	branch_createCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_createCmd.Flags().BoolP("set-default", "d", false, "Also make this branch the default branch, so it is considered the owner of new edits")
	branchCmd.AddCommand(branch_createCmd)
}
