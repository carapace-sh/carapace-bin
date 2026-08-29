package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var refreshRemoteDataCmd = &cobra.Command{
	Use:    "refresh-remote-data",
	Short:  "Trigger a refresh of remote data fetching from the remote, Pull Requests, and CI status",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refreshRemoteDataCmd).Standalone()

	refreshRemoteDataCmd.Flags().Bool("ci", false, "Whether to also refresh CI status from the forge")
	refreshRemoteDataCmd.Flags().Bool("fetch", false, "Whether to also refresh git fetch from the remote")
	refreshRemoteDataCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	refreshRemoteDataCmd.Flags().Bool("pr", false, "Whether to also refresh Pull Requests from the forge")
	refreshRemoteDataCmd.Flags().Bool("updates", false, "Whether to also check for application updates")
	rootCmd.AddCommand(refreshRemoteDataCmd)
}
