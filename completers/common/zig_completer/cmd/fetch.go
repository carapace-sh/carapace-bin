package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch [options] <url>",
	Short:   "Copy a package into global cache and print its hash",
	GroupID: "project",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	rootCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().String("cache-dir", "", "Override path to local cache directory")
	fetchCmd.Flags().Bool("debug-hash", false, "Print verbose hash information to stdout")
	fetchCmd.Flags().String("debug-log", "", "Enable printing debug/info log messages for scope")
	fetchCmd.Flags().BoolP("help", "h", false, "Print help")
	fetchCmd.Flags().String("pkg-dir", "", "Override path to local package directory")
	fetchCmd.Flags().String("save", "", "Add the fetched package to build.zig.zon")
	fetchCmd.Flags().String("save-exact", "", "Add the fetched package to build.zig.zon, storing the URL verbatim")

	fetchCmd.Flags().Lookup("save").NoOptDefVal = " "
	fetchCmd.Flags().Lookup("save-exact").NoOptDefVal = " "

	carapace.Gen(fetchCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir": carapace.ActionDirectories(),
		"pkg-dir":   carapace.ActionDirectories(),
	})

	carapace.Gen(fetchCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
