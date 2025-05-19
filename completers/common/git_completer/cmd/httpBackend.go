package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var httpBackendCmd = &cobra.Command{
	Use:     "http-backend",
	Short:   "Server side implementation of Git over HTTP",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(httpBackendCmd).Standalone()

	rootCmd.AddCommand(httpBackendCmd)
}
