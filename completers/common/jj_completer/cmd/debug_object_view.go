package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_object_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_object_viewCmd).Standalone()

	debug_object_viewCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_object_viewCmd.Flags().String("op", "", "")
	debug_objectCmd.AddCommand(debug_object_viewCmd)
}
