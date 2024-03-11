package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "See information related to packages",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().BoolP("all", "A", false, "Print versions of a package from the whole project")
	infoCmd.Flags().Bool("cache", false, "Print information about the cache entry of a package (path, size, checksum)")
	infoCmd.Flags().Bool("dependents", false, "Print all dependents for each matching package")
	infoCmd.Flags().StringP("extra", "X", "", "An array of requests of extra data provided by plugins")
	infoCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	infoCmd.Flags().Bool("manifest", false, "Print data obtained by looking at the package archive (license, homepage, ...)")
	infoCmd.Flags().Bool("name-only", false, "Only print the name for the matching packages")
	infoCmd.Flags().BoolP("recursive", "R", false, "Print information for all packages, including transitive dependencies")
	infoCmd.Flags().Bool("virtuals", false, "Print each instance of the virtual packages")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).PositionalAnyCompletion(
		yarn.ActionDependencies(),
	)
}
