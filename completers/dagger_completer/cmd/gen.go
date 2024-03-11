package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:    "gen FILE",
	Short:  "Generate CLI reference documentation",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genCmd).Standalone()

	genCmd.Flags().String("frontmatter", "", "Prepend content to beginning of generated file")
	genCmd.Flags().Bool("include-experimental", false, "Show experimental commands in the generated documentation")
	genCmd.Flags().StringP("output", "o", "", "Save to file")
	rootCmd.AddCommand(genCmd)

	carapace.Gen(genCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(genCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
