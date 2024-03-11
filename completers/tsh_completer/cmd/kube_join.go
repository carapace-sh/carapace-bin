package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var kube_joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join an active Kubernetes session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_joinCmd).Standalone()

	kube_joinCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	kube_joinCmd.Flags().StringP("mode", "m", "", "Mode of joining the session, valid modes are observer, moderator and peer.")
	kubeCmd.AddCommand(kube_joinCmd)

	carapace.Gen(kube_joinCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"mode":    carapace.ActionValues("observer", "moderator", "peer"),
	})
}
