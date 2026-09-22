package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codama_convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert an Anchor IDL JSON file (post-0.30 spec) into a Codama IDL rooted at a `rootNode`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codama_convertCmd).Standalone()

	codama_convertCmd.Flags().BoolP("help", "h", false, "Print help")
	codama_convertCmd.Flags().StringP("out", "o", "", "Output file (stdout if not specified)")
	codamaCmd.AddCommand(codama_convertCmd)

	carapace.Gen(codama_convertCmd).FlagCompletion(carapace.ActionMap{
		"out": carapace.ActionFiles(),
	})
}
