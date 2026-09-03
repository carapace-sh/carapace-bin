package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_objectCmd = &cobra.Command{
	Use:   "object",
	Short: "Show information about an operation and its view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_objectCmd).Standalone()

	debug_objectCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_objectCmd)
}
