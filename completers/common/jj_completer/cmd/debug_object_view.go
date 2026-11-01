package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
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

	carapace.Gen(debug_object_viewCmd).FlagCompletion(carapace.ActionMap{
		"op": jj.ActionOperations(100),
	})
}
