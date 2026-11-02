package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().BoolP("check-diff", "d", false, "exit with 1 and show diff if files would be modified by meson format")
	fmtCmd.Flags().BoolP("check-only", "q", false, "silently exit with 1 if files would be modified by meson format")
	fmtCmd.Flags().StringP("configuration", "c", "", "read configuration from meson.format")
	fmtCmd.Flags().BoolP("editor-config", "e", false, "try to read configuration from .editorconfig")
	fmtCmd.Flags().BoolP("inplace", "i", false, "format files in-place")
	fmtCmd.Flags().StringP("output", "o", "", "output file (implies having exactly one input)")
	fmtCmd.Flags().BoolP("recursive", "r", false, "recurse subdirs (requires --check-only, --check-diff or --inplace option)")
	fmtCmd.Flags().String("source-file-path", "", "path to use, when reading from stdin")
	fmtCmd.Flags().BoolP("subprojects", "s", false, "recurse subprojects (requires --recursive)")
	applyCommonFlagCompletion(fmtCmd)
	rootCmd.AddCommand(fmtCmd)
}
