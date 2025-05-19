package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reset_phase_cleanupNodeCmd = &cobra.Command{
	Use:     "cleanup-node",
	Short:   "Run cleanup node.",
	Aliases: []string{"cleanupnode"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reset_phase_cleanupNodeCmd).Standalone()

	reset_phase_cleanupNodeCmd.Flags().String("cert-dir", "", "The path to the directory where the certificates are stored. If specified, clean this directory.")
	reset_phase_cleanupNodeCmd.Flags().Bool("cleanup-tmp-dir", false, "Cleanup the \"/etc/kubernetes/tmp\" directory")
	reset_phase_cleanupNodeCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	reset_phase_cleanupNodeCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	reset_phaseCmd.AddCommand(reset_phase_cleanupNodeCmd)

	carapace.Gen(reset_phase_cleanupNodeCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir":   carapace.ActionDirectories(),
		"cri-socket": carapace.ActionDirectories(),
	})
}
