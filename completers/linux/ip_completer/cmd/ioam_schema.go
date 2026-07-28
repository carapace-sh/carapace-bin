package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioam_schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "manage IOAM schemas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioam_schemaCmd).Standalone()

	ioamCmd.AddCommand(ioam_schemaCmd)

	carapace.Gen(ioam_schemaCmd).PositionalCompletion(
		carapace.ActionValues("show", "add", "del"),
	)
}
