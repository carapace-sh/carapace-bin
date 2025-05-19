package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format *.templ files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().BoolS("help", "help", false, "Print help and exit.")
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
