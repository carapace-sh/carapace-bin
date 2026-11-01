package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "transactional resource manager management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceCmd).Standalone()
	rootCmd.AddCommand(resourceCmd)
}
