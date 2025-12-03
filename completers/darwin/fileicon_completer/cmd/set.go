package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set a custom icon for a file or folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	setCmd.Flags().BoolP("quiet", "q", false, "silence status output")
	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
