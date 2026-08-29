package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_moveCmd = &cobra.Command{
	Use:    "move",
	Short:  "Deprecated: use but move instead",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_moveCmd).Standalone()

	branch_moveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branchCmd.AddCommand(branch_moveCmd)
}
