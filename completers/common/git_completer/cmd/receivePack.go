package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var receivePackCmd = &cobra.Command{
	Use:     "receive-pack",
	Short:   "Receive what is pushed into the repository",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(receivePackCmd).Standalone()

	receivePackCmd.Flags().Bool("http-backend-info-refs", false, "Used by git-http-backend to serve up requests")
	receivePackCmd.Flags().Bool("skip-connectivity-check", false, "Bypasses the connectivity checks")
	rootCmd.AddCommand(receivePackCmd)

	carapace.Gen(receivePackCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
