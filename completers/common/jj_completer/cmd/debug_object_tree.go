package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_object_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_object_treeCmd).Standalone()

	debug_object_treeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_object_treeCmd.Flags().StringP("revision", "r", "", "")
	debug_objectCmd.AddCommand(debug_object_treeCmd)

	carapace.Gen(debug_object_treeCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
