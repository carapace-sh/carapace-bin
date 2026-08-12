package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var file_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write data to a specific location",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_writeCmd).Standalone()

	file_writeCmd.Flags().String("address", "", "Address of a blob")
	file_writeCmd.Flags().BoolP("help", "h", false, "Print help")
	file_writeCmd.Flags().String("output", "", "Path to a destination")
	file_writeCmd.Flags().String("path", "", "Path to a file")
	file_writeCmd.Flags().String("revision", "", "Revision specifier")
	file_writeCmd.MarkFlagRequired("output")
	fileCmd.AddCommand(file_writeCmd)

	carapace.Gen(file_writeCmd).FlagCompletion(carapace.ActionMap{
		"address":  carapace.ActionValues(),
		"output":   carapace.ActionFiles(),
		"path":     carapace.ActionFiles(),
		"revision": action.ActionRevisions(file_writeCmd),
	})
}
