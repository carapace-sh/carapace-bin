package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_cdnCmd = &cobra.Command{
	Use:   "cdn",
	Short: "Display commands that manage CDNs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_cdnCmd).Standalone()
	computeCmd.AddCommand(compute_cdnCmd)
}
