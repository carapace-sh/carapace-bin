package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var debug_object_symlinkCmd = &cobra.Command{
	Use:   "symlink",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_object_symlinkCmd).Standalone()

	debug_object_symlinkCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_object_symlinkCmd.Flags().StringP("revision", "r", "", "")
	debug_objectCmd.AddCommand(debug_object_symlinkCmd)

	carapace.Gen(debug_object_symlinkCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevsets(jj.RevOpts{}.Default()),
	})

	carapace.Gen(debug_object_symlinkCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
