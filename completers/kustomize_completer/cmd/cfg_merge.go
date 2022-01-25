package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "[Alpha] Merge Resource configuration files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_mergeCmd).Standalone()
	cfg_mergeCmd.Flags().Bool("invert-order", false, "if true, merge Resources in the reverse order")
	cfgCmd.AddCommand(cfg_mergeCmd)
}
