package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_createCmd).Standalone()

	branch_createCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_createCmd.Flags().String("id", "", "Optional explicit branch ID (hex-encoded 16-byte identifier)")
	branchCmd.AddCommand(branch_createCmd)
}
