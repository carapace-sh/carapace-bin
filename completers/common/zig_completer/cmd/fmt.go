package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Reformat Zig source into canonical form",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().Bool("check", false, "List non-conforming files and exit with an error if the list is non-empty")
	fmtCmd.Flags().Bool("ast-check", false, "Run zig ast-check on every file")
	fmtCmd.Flags().Bool("color", false, "Enable or disable colored error messages")
	fmtCmd.Flags().Bool("stdin", false, "Format code from stdin; output to stdout")
	fmtCmd.Flags().StringP("exclude", "", "", "Exclude file or directory from formatting")
	fmtCmd.Flags().BoolP("help", "h", false, "Print help")

	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".zig"),
	)
}
