package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioam_namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "manage IOAM namespaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioam_namespaceCmd).Standalone()

	ioamCmd.AddCommand(ioam_namespaceCmd)

	carapace.Gen(ioam_namespaceCmd).PositionalCompletion(
		carapace.ActionValues("show", "add", "del", "set"),
	)
}
