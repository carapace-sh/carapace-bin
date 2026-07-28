package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_statedirCmd = &cobra.Command{
	Use:   "statedir",
	Short: "Print the location of the state directory (if any)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_statedirCmd).Standalone()

	debugCmd.AddCommand(debug_statedirCmd)
}
