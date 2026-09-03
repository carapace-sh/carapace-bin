package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_initSimpleCmd = &cobra.Command{
	Use:   "init-simple",
	Short: "Create a new repo in the given directory using the proof-of-concept simple backend",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_initSimpleCmd).Standalone()

	debug_initSimpleCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_initSimpleCmd)

	carapace.Gen(debug_initSimpleCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}