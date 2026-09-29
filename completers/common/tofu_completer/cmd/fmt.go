package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt [options] [target...]",
	Short: "Reformat your configuration in the standard style",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().BoolS("check", "check", false, "Check if the input is formatted. Exit status will be 0 if all input is properly formatted and nonzero otherwise.")
	fmtCmd.Flags().BoolS("diff", "diff", false, "Display diffs of formatting changes")
	fmtCmd.Flags().BoolS("list", "list", false, "Don't list files whose formatting differs")
	fmtCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color")
	fmtCmd.Flags().BoolS("recursive", "recursive", false, "Also process files in subdirectories. By default, fmt only processes the current directory")
	fmtCmd.Flags().BoolS("write", "write", false, "Don't write to source files")
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
