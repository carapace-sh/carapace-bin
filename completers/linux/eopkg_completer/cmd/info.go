package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <package-name>",
	Short: "show information about the given package name or package file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().StringP("component", "c", "", "show information about a component instead of a package")
	infoCmd.Flags().BoolP("files", "f", false, "show a list of the package's files if available")
	infoCmd.Flags().StringP("files-path", "F", "", "only show the files, and no other information about the package")
	infoCmd.Flags().BoolP("short", "s", false, "compact information about each package")
	infoCmd.Flags().Bool("xml", false, "emit the original XML metadata for the package")

	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles("eopkg"),
			eopkg.ActionPackageSearch(),
		).ToA().FilterArgs(),
	)
}
