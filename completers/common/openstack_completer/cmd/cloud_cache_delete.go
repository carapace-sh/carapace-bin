package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud_cache_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the authentication cached for a cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_cache_deleteCmd).Standalone()

	cloud_cacheCmd.AddCommand(cloud_cache_deleteCmd)
}
