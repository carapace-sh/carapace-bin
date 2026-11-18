package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_help_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import every thread the agents have on this machine, including those capture leaves alone: the ones nobody has gone on with since it was turned on",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_help_importCmd).Standalone()

	threads_helpCmd.AddCommand(threads_help_importCmd)
}
