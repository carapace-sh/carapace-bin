package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var credentialCacheCmd = &cobra.Command{
	Use:     "credential-cache",
	Short:   "Helper to temporarily store passwords in memory",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(credentialCacheCmd).Standalone()

	credentialCacheCmd.Flags().String("socket", "", "path of cache-daemon socket")
	credentialCacheCmd.Flags().String("timeout", "", "number of seconds to cache credentials")
	rootCmd.AddCommand(credentialCacheCmd)

	carapace.Gen(credentialCacheCmd).FlagCompletion(carapace.ActionMap{
		"socket": carapace.ActionFiles(),
	})
}
