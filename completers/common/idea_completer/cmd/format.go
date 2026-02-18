package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Apply code style formatting to the specified files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()

	formatCmd.Flags().BoolN("R", "r", false, "Process specified directories recursively")
	formatCmd.Flags().BoolS("allowDefaults", "allowDefaults", false, "Use the default code style settings when the code style is not defined for a file")
	formatCmd.Flags().StringS("charset", "charset", "", "Preserve encoding and enforce the charset for reading and writing source files")
	formatCmd.Flags().BoolN("dry", "d", false, "Run the formatter in the validation mode")
	formatCmd.Flags().BoolS("h", "h", false, "Show help for format command")
	formatCmd.Flags().StringN("mask", "m", "", "Specify a comma-separated list of file masks that define the files to be processed")
	formatCmd.Flags().StringN("settings", "s", "", "Specify the code style settings file to use for formatting")

	rootCmd.AddCommand(formatCmd)

	carapace.Gen(formatCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
