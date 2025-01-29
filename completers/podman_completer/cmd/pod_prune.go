package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_pruneCmd = &cobra.Command{
	Use:   "prune [options]",
	Short: "Remove all stopped pods and their containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_pruneCmd).Standalone()

	pod_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation.  The default is false")
	podCmd.AddCommand(pod_pruneCmd)
}
