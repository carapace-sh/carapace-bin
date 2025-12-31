package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sshCacheCmd = &cobra.Command{
	Use:   "+ssh-cache",
	Short: "Manage the SSH terminfo cache for automatic remote host setup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshCacheCmd).Standalone()

	sshCacheCmd.Flags().String("add", "", "Manually add host to cache")
	sshCacheCmd.Flags().Bool("clear", false, "Clear entire cache")
	sshCacheCmd.Flags().String("expire-days", "", "Set custom expiration period")
	sshCacheCmd.Flags().String("host", "", "Check if host is cached")
	sshCacheCmd.Flags().String("remove", "", "Remove host from cache")
	rootCmd.AddCommand(sshCacheCmd)

	sshCacheCmd.MarkFlagsMutuallyExclusive("clear", "add", "remove", "host")
}
