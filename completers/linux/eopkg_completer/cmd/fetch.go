package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch <name>",
	Aliases: []string{"fc"},
	Short:   "download the package file for the named package, into the current working directory",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().StringP("output-dir", "o", "", "override the output directory for the .eopkg")

	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).PositionalAnyCompletion(
		eopkg.ActionPackageSearch().FilterArgs(),
	)
}
