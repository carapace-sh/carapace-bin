package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-dgraph [OPTIONS] [MODE] <pkgname>",
	Short: "XBPS utility to generate package dot graphs",
	Long:  "https://man.voidlinux.org/xbps-dgraph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config", "C", "", "Path to confdir (xbps.d)")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode shown to stderr")
	rootCmd.Flags().BoolP("fulldeptree", "f", false, "Generate a dependency graph")
	rootCmd.Flags().BoolP("gen-config", "g", false, "Generate a configuration file")
	rootCmd.Flags().StringP("graph-config", "c", "", "Path to the graph configuration file")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().BoolP("memory-sync", "M", false, "Remote repository data is fetched and stored in memory, ignoring on-disk repodata archives.")
	rootCmd.Flags().BoolP("metadata", "m", false, "Generate a metadata graph (default mode)")
	rootCmd.Flags().BoolP("repository", "R", false, "Enable repository mode. This mode explicitly looks for packages in repositories.")
	rootCmd.Flags().StringP("rootdir", "r", "", "Full path to rootdir")

	rootCmd.MarkFlagsMutuallyExclusive("gen-config", "fulldeptree", "metadata")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":       carapace.ActionDirectories(),
		"graph-config": carapace.ActionFiles(),
		"rootdir":      carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(xbps.ActionPackages().FilterArgs())
}
