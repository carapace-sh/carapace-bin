package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Copy a package into the global cache and print its hash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().StringP("global-cache-dir", "", "", "Override path to global Zig cache directory")
	fetchCmd.Flags().Bool("debug-hash", false, "Print verbose hash information to stdout")
	fetchCmd.Flags().Bool("save", false, "Add the fetched package to build.zig.zon")
	fetchCmd.Flags().String("save-exact", "", "Add the fetched package to build.zig.zon, pinned to the precise URL")
	fetchCmd.Flags().BoolP("help", "h", false, "Print help")

	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).FlagCompletion(carapace.ActionMap{
		"global-cache-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(fetchCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
