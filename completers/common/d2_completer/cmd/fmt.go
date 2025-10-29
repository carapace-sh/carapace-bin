package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt file.d2 ...",
	Short: "Format all passed files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".d2"),
	)
}
