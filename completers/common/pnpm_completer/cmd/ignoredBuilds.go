package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ignoredBuildsCmd = &cobra.Command{
	Use:   "ignored-builds",
	Short: "Print the list of packages with blocked build scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ignoredBuildsCmd).Standalone()

	ignoredBuildsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(ignoredBuildsCmd)
}
