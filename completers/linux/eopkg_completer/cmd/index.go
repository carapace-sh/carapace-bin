package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:     "index",
	Short:   "Index eopkg packages",
	Aliases: []string{"ix"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(indexCmd).Standalone()

	indexCmd.Flags().BoolP("absolute-urls", "a", false, "Store absolute links for index file")
	indexCmd.Flags().StringP("output", "o", "", "Index output file")
	indexCmd.Flags().Bool("skip-sources", false, "Do not include sources in index")
	indexCmd.Flags().Bool("skip-signing", false, "Do not sign the index")

	rootCmd.AddCommand(indexCmd)

	carapace.Gen(indexCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(indexCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
