package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "Fetch dependencies of a package from the network",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("fetch"),
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().BoolP("help", "h", false, "Print help")
	fetchCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	fetchCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	fetchCmd.Flags().StringSlice("target", nil, "Fetch dependencies for the target triple")
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).FlagCompletion(carapace.ActionMap{
		"lockfile-path": carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
	})
}
