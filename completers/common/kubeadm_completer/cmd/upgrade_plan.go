package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kubeadm_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubeadm"
	"github.com/spf13/cobra"
)

var upgrade_planCmd = &cobra.Command{
	Use:   "plan [version] [flags]",
	Short: "Check which versions are available to upgrade to and validate whether your current cluster is upgradeable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_planCmd).Standalone()

	upgrade_planCmd.Flags().Bool("allow-experimental-upgrades", false, "Show unstable versions of Kubernetes as an upgrade alternative and allow upgrading to an alpha/beta/release candidate versions of Kubernetes.")
	upgrade_planCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	upgrade_planCmd.Flags().Bool("allow-release-candidate-upgrades", false, "Show release candidate versions of Kubernetes as an upgrade alternative and allow upgrading to a release candidate versions of Kubernetes.")
	upgrade_planCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_planCmd.Flags().Bool("etcd-upgrade", false, "Perform the upgrade of etcd.")
	upgrade_planCmd.Flags().StringSlice("ignore-preflight-errors", nil, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	upgrade_planCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_planCmd.Flags().StringP("output", "o", "", "Output format. One of: text|json|yaml|kyaml|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file.")
	upgrade_planCmd.Flags().Bool("print-config", false, "Specifies whether the configuration file that will be used in the upgrade should be printed or not.")
	upgrade_planCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	upgradeCmd.AddCommand(upgrade_planCmd)

	carapace.Gen(upgrade_planCmd).FlagCompletion(carapace.ActionMap{
		"config":                  carapace.ActionFiles(),
		"ignore-preflight-errors": action.ActionChecks().UniqueList(","),
		"kubeconfig":              carapace.ActionFiles(),
		"output":                  kubeadm.ActionOutputFormats(),
	})
}
