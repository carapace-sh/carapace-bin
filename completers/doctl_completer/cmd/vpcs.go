package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var vpcsCmd = &cobra.Command{
	Use:   "vpcs",
	Short: "Display commands that manage VPCs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcsCmd).Standalone()
	rootCmd.AddCommand(vpcsCmd)
}
