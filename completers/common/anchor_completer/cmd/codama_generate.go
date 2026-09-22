package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codama_generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Convert an Anchor IDL and run Codama renderers to produce client libraries in one or more languages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codama_generateCmd).Standalone()

	codama_generateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	codama_generateCmd.Flags().StringSliceP("language", "l", nil, "Languages to generate clients for. Repeat the flag or comma- separate values: `-l js,go -l rust`")
	codama_generateCmd.Flags().StringP("path", "p", "clients", "Base output directory; per-language clients are written to `<path>/<language>`")
	codama_generateCmd.MarkFlagRequired("language")
	codamaCmd.AddCommand(codama_generateCmd)

	carapace.Gen(codama_generateCmd).FlagCompletion(carapace.ActionMap{
		"language": carapace.ActionValues("js", "js-umi", "rust", "go"),
		"path":     carapace.ActionFiles(),
	})
}
