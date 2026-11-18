package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Turn thread capture on (short for `atuin config enable threads`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_enableCmd).Standalone()

	threads_enableCmd.Flags().BoolP("help", "h", false, "Print help")
	threadsCmd.AddCommand(threads_enableCmd)
}
