package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Reformat your configuration in the standard style",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().Bool("check", false, "Check if the input is formatted")
	fmtCmd.Flags().Bool("diff", false, "Display diffs of formatting changes")
	fmtCmd.Flags().String("list", "", "Don't list files whose formatting differs")
	fmtCmd.Flags().Bool("no-color", false, "If specified, output won't contain any color")
	fmtCmd.Flags().Bool("recursive", false, "Also process files in subdirectories")
	fmtCmd.Flags().String("write", "", "Don't write to source files")
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).FlagCompletion(carapace.ActionMap{
		"list":  action.ActionBool(),
		"write": action.ActionBool(),
	})

	carapace.Gen(fmtCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
