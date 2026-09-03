package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Show commit index stats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_indexCmd).Standalone()

	debug_indexCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_indexCmd)
}
