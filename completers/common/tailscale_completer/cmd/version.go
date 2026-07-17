package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print Tailscale version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().Bool("daemon", false, "also print local node's daemon version")
	versionCmd.Flags().Bool("json", false, "output in JSON format")
	versionCmd.Flags().String("track", "", "which track to check for updates: stable, release-candidate, or unstable")
	versionCmd.Flags().Bool("upstream", false, "fetch and print the latest upstream release version")
	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).FlagCompletion(carapace.ActionMap{
		"track": carapace.ActionValues("stable", "release-candidate", "unstable"),
	})
}
