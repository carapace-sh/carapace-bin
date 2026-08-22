package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var delegation_createSessionCmd = &cobra.Command{
	Use:   "create-session",
	Short: "Create a delegation session, allowing a bot or workload to temporarily act on your behalf.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delegation_createSessionCmd).Standalone()

	delegation_createSessionCmd.Flags().Bool("allow-all", false, "Allow access to all resources, including destructive administrative actions. Mutually exclusive with the other --allow-* flags.")
	delegation_createSessionCmd.Flags().String("allow-app", "", "Allow access to an application.")
	delegation_createSessionCmd.Flags().String("allow-db", "", "Allow access to a database.")
	delegation_createSessionCmd.Flags().String("allow-git-server", "", "Allow access to a Git server.")
	delegation_createSessionCmd.Flags().String("allow-kube-cluster", "", "Allow access to a Kubernetes cluster.")
	delegation_createSessionCmd.Flags().String("allow-node", "", "Allow access to an SSH node.")
	delegation_createSessionCmd.Flags().String("allow-windows-desktop", "", "Allow access to a Windows desktop.")
	delegation_createSessionCmd.Flags().String("bot", "", "Name of a bot allowed to use the delegation session. Repeat to allow multiple bots.")
	delegation_createSessionCmd.Flags().Bool("no-allow-all", false, "Allow access to all resources, including destructive administrative actions. Mutually exclusive with the other --allow-* flags.")
	delegation_createSessionCmd.Flags().String("session-ttl", "", "How long the delegation session should remain valid.")
	delegation_createSessionCmd.Flag("no-allow-all").Hidden = true
	delegationCmd.AddCommand(delegation_createSessionCmd)
}
