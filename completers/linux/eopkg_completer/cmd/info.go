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
	infoCmd.Flags().BoolP("files-path", "F", false, "show only paths")
	infoCmd.Flags().StringP("repo", "r", "", "resolve package against specified repository")
	infoCmd.Flags().BoolP("short", "s", false, "do not show details")
	infoCmd.Flags().String("sort-by", "", "sort files by name (path), size or type")
	infoCmd.Flags().Bool("xml", false, "emit the original XML metadata for the package")

	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"component": eopkg.ActionComponents(),
		"repo":      eopkg.ActionRepositories(),
		"sort-by":   carapace.ActionValues("path", "size", "type"),
	})

	carapace.Gen(infoCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles("eopkg"),
			eopkg.ActionPackageSearch(),
		).ToA().FilterArgs(),
	)
}
