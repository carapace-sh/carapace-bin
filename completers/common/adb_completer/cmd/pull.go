package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull [-a] [-z ALGORITHM] [-Z] REMOTE... LOCAL",
	Short: "copy files/dirs from device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().BoolS("Z", "Z", false, "disable compression")
	pullCmd.Flags().BoolS("a", "a", false, "preserve file timestamp and mode")
	pullCmd.Flags().StringS("z", "z", "", "enable compression with a specified algorithm (any, none, brotli)")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).FlagCompletion(carapace.ActionMap{
		"z": carapace.ActionValues("any", "none", "brotli"),
	})

	carapace.Gen(pullCmd).PositionalCompletion(
		adb.ActionFiles(),
	)

	carapace.Gen(pullCmd).PositionalAnyCompletion(
		carapace.Batch(
			adb.ActionFiles().FilterArgs(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
