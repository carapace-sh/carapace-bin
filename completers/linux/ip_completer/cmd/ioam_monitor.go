package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioam_monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "display IOAM data received",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioam_monitorCmd).Standalone()

	ioamCmd.AddCommand(ioam_monitorCmd)
}
