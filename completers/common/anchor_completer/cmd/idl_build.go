package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/anchor"
	"github.com/spf13/cobra"
)

var idl_buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Generates the IDL for the program using the compilation method",
	Aliases: []string{"b"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_buildCmd).Standalone()

	idl_buildCmd.Flags().BoolP("help", "h", false, "Print help")
	idl_buildCmd.Flags().Bool("no-docs", false, "Suppress doc strings in output")
	idl_buildCmd.Flags().StringP("out", "o", "", "Output file for the IDL (stdout if not specified)")
	idl_buildCmd.Flags().StringP("out-ts", "t", "", "Output file for the TypeScript IDL")
	idl_buildCmd.Flags().StringP("program-name", "p", "", "")
	idl_buildCmd.Flags().Bool("skip-lint", false, "Do not check for safety comments")
	idlCmd.AddCommand(idl_buildCmd)

	carapace.Gen(idl_buildCmd).FlagCompletion(carapace.ActionMap{
		"out":          carapace.ActionFiles(),
		"out-ts":       carapace.ActionFiles(".json"),
		"program-name": anchor.ActionPrograms(),
	})
}
