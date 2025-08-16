package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/typst_completer/cmd/common"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compiles an input file into a supported output format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compileCmd).Standalone()

	compileCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")

	common.AddPdfFlags(compileCmd)

	rootCmd.AddCommand(compileCmd)

	carapace.Gen(compileCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)

}
