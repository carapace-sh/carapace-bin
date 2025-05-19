package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_diffCmd = &cobra.Command{
	Use:   "diff [version]",
	Short: "Show what differences would be applied to existing static pod manifests. See also: kubeadm upgrade apply --dry-run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_diffCmd).Standalone()

	upgrade_diffCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_diffCmd.Flags().StringP("context-lines", "c", "", "How many lines of context in the diff")
	upgrade_diffCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgradeCmd.AddCommand(upgrade_diffCmd)

	carapace.Gen(upgrade_diffCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
