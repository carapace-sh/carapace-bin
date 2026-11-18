package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_threads_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Turn thread capture on (short for `atuin config enable threads`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_threads_enableCmd).Standalone()

	help_threadsCmd.AddCommand(help_threads_enableCmd)
}
