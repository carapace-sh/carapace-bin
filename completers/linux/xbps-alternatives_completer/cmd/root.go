package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-alternatives [OPTIONS] [MODE]",
	Short: "XBPS utility to handle alternatives",
	Long:  "https://man.voidlinux.org/xbps-alternatives",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("R", "R", false, "Enable repository mode. This mode explicitly looks for packages in repositories")
	rootCmd.Flags().StringP("config", "C", "", "Path to confdir (xbps.d)")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode shown to stderr")
	rootCmd.Flags().StringP("group", "g", "", "Group of alternatives to match")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().BoolP("ignore-conf-repos", "i", false, "Ignore repositories defined in xbps.d")
	rootCmd.Flags().StringP("list", "l", "", "List all alternatives or from PKG")
	rootCmd.Flags().StringArray("repository", nil, "Enable repository mode and add repository to the top of the list. This option can be specified multiple times")
	rootCmd.Flags().StringP("rootdir", "r", "", "Full path to rootdir")
	rootCmd.Flags().StringP("set", "s", "", "Set alternatives for PKG")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose messages")
	rootCmd.Flags().BoolP("version", "V", false, "Show XBPS version")

	rootCmd.MarkFlagsMutuallyExclusive("list", "set")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":  carapace.ActionDirectories(),
		"list":    xbps.ActionPackageSearch(),
		"rootdir": carapace.ActionDirectories(),
		"set":     xbps.ActionPackageSearch(),
	})
}
