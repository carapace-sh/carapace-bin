package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idl_convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert legacy IDLs (pre Anchor 0.30) to the new IDL spec",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_convertCmd).Standalone()

	idl_convertCmd.Flags().BoolP("help", "h", false, "Print help")
	idl_convertCmd.Flags().StringP("out", "o", "", "Output file for the IDL (stdout if not specified)")
	idl_convertCmd.Flags().StringP("program-id", "p", "", "Program id to initialize IDL for. If not provided, discovers program ID from IDL")
	idl_convertCmd.Flags().Bool("to-legacy", false, "Convert a current-spec IDL back to the legacy (pre Anchor v0.30) format. Without this flag the converter runs in the default direction (legacy -> current)")
	idlCmd.AddCommand(idl_convertCmd)

	carapace.Gen(idl_convertCmd).FlagCompletion(carapace.ActionMap{
		"out": carapace.ActionFiles(),
	})

	carapace.Gen(idl_convertCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
