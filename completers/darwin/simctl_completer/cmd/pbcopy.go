package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var pbcopyCmd = &cobra.Command{
	Use:   "pbcopy",
	Short: "Copy standard input onto the device pasteboard",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pbcopyCmd).Standalone()
	rootCmd.AddCommand(pbcopyCmd)
	carapace.Gen(pbcopyCmd).PositionalCompletion(simctl.ActionDevices())
}
