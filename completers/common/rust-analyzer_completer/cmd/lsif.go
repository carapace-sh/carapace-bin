package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lsifCmd = &cobra.Command{
	Use:   "lsif",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsifCmd).Standalone()

	rootCmd.AddCommand(lsifCmd)

	carapace.Gen(lsifCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
