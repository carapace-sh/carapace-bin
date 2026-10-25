package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join the active SSH or Kubernetes session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(joinCmd).Standalone()

	joinCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	joinCmd.Flags().StringP("mode", "m", "observer", "Mode of joining the session, valid modes are observer, moderator and peer.")
	rootCmd.AddCommand(joinCmd)

	carapace.Gen(joinCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"mode":    carapace.ActionValues("moderator", "observer", "peer"),
	})

	carapace.Gen(joinCmd).PositionalCompletion(
		tsh.ActionSessions("ssh,k8s"),
	)
}
