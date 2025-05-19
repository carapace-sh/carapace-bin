package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgrade_nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Upgrade commands for a node in the cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_nodeCmd).Standalone()

	upgrade_nodeCmd.Flags().Bool("certificate-renewal", false, "Perform the renewal of certificates used by component changed during upgrades.")
	upgrade_nodeCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_nodeCmd.Flags().Bool("dry-run", false, "Do not change any state, just output the actions that would be performed.")
	upgrade_nodeCmd.Flags().Bool("etcd-upgrade", false, "Perform the upgrade of etcd.")
	upgrade_nodeCmd.Flags().StringSlice("ignore-preflight-errors", nil, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	upgrade_nodeCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_nodeCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\". For example, \"kube-apiserver0+merge.yaml\" or just \"etcd.json\". \"target\" can be one of \"kube-apiserver\", \"kube-controller-manager\", \"kube-scheduler\", \"etcd\", \"kubeletconfiguration\", \"corednsdeployment\". \"patchtype\" can be one of \"strategic\", \"merge\" or \"json\" and they match the patch formats supported by kubectl. The default \"patchtype\" is \"strategic\". \"extension\" must be either \"json\" or \"yaml\". \"suffix\" is an optional string that can be used to determine which patches are applied first alpha-numerically.")
	upgrade_nodeCmd.Flags().StringSlice("skip-phases", nil, "List of phases to be skipped")
	upgradeCmd.AddCommand(upgrade_nodeCmd)

	carapace.Gen(upgrade_nodeCmd).FlagCompletion(carapace.ActionMap{
		"ignore-preflight-errors": action.ActionChecks().UniqueList(","),
		"kubeconfig":              carapace.ActionFiles(),
		"patches":                 carapace.ActionDirectories(),
		"skip-phases": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionPhases().Invoke(c).Filter(c.Parts...).ToMultiPartsA("/").NoSpace() // TODO use UniqueList("/")
		}),
	})
}
