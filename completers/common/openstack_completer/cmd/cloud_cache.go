package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud_cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_cacheCmd).Standalone()

	cloudCmd.AddCommand(cloud_cacheCmd)
}
