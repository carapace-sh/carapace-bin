package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sshKeyCmd = &cobra.Command{
	Use:     "ssh-key",
	Short:   "Retrieve the ssh identity key path of the specified node",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKeyCmd).Standalone()

	sshKeyCmd.Flags().StringP("node", "n", "", "The node to get ssh-key path. Defaults to the primary control plane.")
	rootCmd.AddCommand(sshKeyCmd)

	carapace.Gen(sshKeyCmd).FlagCompletion(carapace.ActionMap{
		"node": action.ActionNodes(),
	})
}
