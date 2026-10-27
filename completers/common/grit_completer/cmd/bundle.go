package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundleCmd = &cobra.Command{
	Use:     "bundle",
	Short:   "Create, verify, or list git bundle files",
	GroupID: "remotes",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundleCmd).Standalone()

	bundleCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(bundleCmd)
}
