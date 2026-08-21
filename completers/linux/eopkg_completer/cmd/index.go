package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:     "index <directory>",
	Aliases: []string{"ix"},
	Short:   "produce an eopkg-index repository in the given directory",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(indexCmd).Standalone()

	indexCmd.Flags().BoolP("absolute-urls", "a", false, "use absolute URLs in the index instead of relative ones")
	indexCmd.Flags().String("compression-types", "", "comma separated list of compression types to use when producing the index")
	indexCmd.Flags().StringP("output", "o", "", "override path to the output file")
	indexCmd.Flags().Bool("skip-signing", false, "do not attempt to GPG sign the index")
	indexCmd.Flags().Bool("skip-sources", false, "do not include pspec.xml legacy format eopkg definitions in the index")

	rootCmd.AddCommand(indexCmd)

	carapace.Gen(indexCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
