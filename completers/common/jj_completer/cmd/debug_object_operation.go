package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_object_operationCmd = &cobra.Command{
	Use:   "operation",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_object_operationCmd).Standalone()

	debug_object_operationCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_objectCmd.AddCommand(debug_object_operationCmd)
}