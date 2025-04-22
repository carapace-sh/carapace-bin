package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var kube_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to a Kubernetes cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_loginCmd).Standalone()

	kube_loginCmd.Flags().Bool("all", false, "Generate a kubeconfig with every cluster the user has access to.")
	kube_loginCmd.Flags().String("as", "", "Configure custom Kubernetes user impersonation.")
	kube_loginCmd.Flags().String("as-groups", "", "Configure custom Kubernetes group impersonation.")
	kube_loginCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	kube_loginCmd.Flags().StringP("kube-namespace", "n", "", "Configure the default Kubernetes namespace.")
	kube_loginCmd.Flags().Bool("no-all", false, "Generate a kubeconfig with every cluster the user has access to.")
	kube_loginCmd.Flags().String("set-context-name", "", "Define a custom context name.")
	kube_loginCmd.Flag("no-all").Hidden = true
	kubeCmd.AddCommand(kube_loginCmd)

	carapace.Gen(kube_loginCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
	})

	carapace.Gen(kube_loginCmd).PositionalCompletion(
		tsh.ActionKubeClusters(),
	)
}
