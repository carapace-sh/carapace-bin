package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_setBaseCmd = &cobra.Command{
	Use:   "set-base",
	Short: "Switch to the GitButler workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_setBaseCmd).Standalone()

	branch_setBaseCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_setBaseCmd)
}
