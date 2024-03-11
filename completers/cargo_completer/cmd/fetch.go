package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("fetch"),
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().BoolP("help", "h", false, "Print help")
	fetchCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	fetchCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	fetchCmd.Flags().StringSlice("target", []string{}, "Fetch dependencies for the target triple")
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
