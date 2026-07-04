package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var reformatCmd = &cobra.Command{
	Use:   "reformat",
	Short: "Reformat an existing volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reformatCmd).Standalone()
	rootCmd.AddCommand(reformatCmd)
	carapace.Gen(reformatCmd).PositionalCompletion(
		carapace.ActionValues("APFS", "HFS+", "ExFAT", "MS-DOS"),
		fs.ActionBlockDevices(),
	)
}
