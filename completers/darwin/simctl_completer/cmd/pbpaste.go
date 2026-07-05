package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var pbpasteCmd = &cobra.Command{
	Use:   "pbpaste",
	Short: "Print the contents of the device's pasteboard to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pbpasteCmd).Standalone()
	rootCmd.AddCommand(pbpasteCmd)
	carapace.Gen(pbpasteCmd).PositionalCompletion(simctl.ActionDevices())
}
