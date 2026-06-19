package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdistCmd = &cobra.Command{
	Use:   "sdist",
	Short: "Build only a source distribution (sdist) without compiling",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdistCmd).Standalone()

	sdistCmd.Flags().StringP("manifest-path", "m", "", "The path to the Cargo.toml")
	sdistCmd.Flags().StringP("out", "o", "", "The directory to store the built wheels in")
	sdistCmd.Flags().CountP("verbose", "v", "Use verbose output")
	rootCmd.AddCommand(sdistCmd)

	carapace.Gen(sdistCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"out":           carapace.ActionDirectories(),
	})
}
