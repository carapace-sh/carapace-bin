package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format meson source file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()

	formatCmd.Flags().BoolP("check-diff", "d", false, "exit with 1 and show diff if files would be modified by meson format")
	formatCmd.Flags().BoolP("check-only", "q", false, "silently exit with 1 if files would be modified by meson format")
	formatCmd.Flags().StringP("configuration", "c", "", "read configuration from meson.format")
	formatCmd.Flags().BoolP("editor-config", "e", false, "try to read configuration from .editorconfig")
	formatCmd.Flags().BoolP("inplace", "i", false, "format files in-place")
	formatCmd.Flags().StringP("output", "o", "", "output file (implies having exactly one input)")
	formatCmd.Flags().BoolP("recursive", "r", false, "recurse subdirs (requires --check-only, --check-diff or --inplace option)")
	formatCmd.Flags().String("source-file-path", "", "path to use, when reading from stdin")
	formatCmd.Flags().BoolP("subprojects", "s", false, "recurse subprojects (requires --recursive)")
	applyCommonFlagCompletion(formatCmd)
	rootCmd.AddCommand(formatCmd)
}
