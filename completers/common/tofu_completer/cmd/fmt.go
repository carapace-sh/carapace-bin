package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt [options] [paths...]",
	Short: "Reformat your configuration in the standard style",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().BoolS("check", "check", false, "Check if the input is formatted. Exit status will be 0 if all input is properly formatted and nonzero otherwise.")
	fmtCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	fmtCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	fmtCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	fmtCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	fmtCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	fmtCmd.Flags().BoolS("diff", "diff", false, "Display diffs of formatting changes")
	fmtCmd.Flags().BoolS("list", "list", false, "Don't list files whose formatting differs")
	fmtCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color")
	fmtCmd.Flags().BoolS("recursive", "recursive", false, "Also process files in subdirectories. By default, fmt only processes the current directory")
	fmtCmd.Flags().BoolS("write", "write", false, "Don't write to source files")
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
	})

	carapace.Gen(fmtCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
