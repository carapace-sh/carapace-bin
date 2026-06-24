package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bisect_nextCmd = &cobra.Command{
	Use:   "next",
	Short: "find next bisection commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_nextCmd).Standalone()

	bisectCmd.AddCommand(bisect_nextCmd)
}
