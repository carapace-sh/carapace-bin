package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translateCCmd = &cobra.Command{
	Use:   "translate-c",
	Short: "Convert C code to Zig code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translateCCmd).Standalone()

	translateCCmd.Flags().StringP("target", "", "", "<arch><sub>-<os>-<abi>")
	translateCCmd.Flags().StringP("cflags", "", "", "Set extra flags for the C compiler")
	translateCCmd.Flags().BoolP("help", "h", false, "Print help")

	rootCmd.AddCommand(translateCCmd)

	carapace.Gen(translateCCmd).PositionalCompletion(
		carapace.ActionFiles(".c", ".h"),
	)
}
