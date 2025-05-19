package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Performs a best effort revert of changes made to this host by 'kubeadm init' or 'kubeadm join'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetCmd).Standalone()

	resetCmd.Flags().String("cert-dir", "", "The path to the directory where the certificates are stored. If specified, clean this directory.")
	resetCmd.Flags().Bool("cleanup-tmp-dir", false, "Cleanup the \"/etc/kubernetes/tmp\" directory")
	resetCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	resetCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	resetCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	resetCmd.Flags().BoolP("force", "f", false, "Reset the node without prompting for confirmation.")
	resetCmd.Flags().StringSlice("ignore-preflight-errors", nil, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	resetCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	resetCmd.Flags().StringSlice("skip-phases", nil, "List of phases to be skipped")
	rootCmd.AddCommand(resetCmd)

	carapace.Gen(resetCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":                carapace.ActionDirectories(),
		"cri-socket":              carapace.ActionFiles(),
		"ignore-preflight-errors": action.ActionChecks().UniqueList(","),
		"kubeconfig":              carapace.ActionFiles(),
		"skip-phases":             action.ActionPhases().UniqueList(","),
	})
}
