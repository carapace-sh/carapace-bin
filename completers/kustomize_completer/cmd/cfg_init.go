package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_initCmd = &cobra.Command{
	Use:   "init",
	Short: "[Alpha] Initialize a directory with a Krmfile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_initCmd).Standalone()
	cfgCmd.AddCommand(cfg_initCmd)
}
