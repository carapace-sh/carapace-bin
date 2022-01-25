package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_merge3Cmd = &cobra.Command{
	Use:   "merge3",
	Short: "[Alpha] Merge diff of Resource configuration files into a destination (3-way)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_merge3Cmd).Standalone()
	cfg_merge3Cmd.Flags().String("ancestor", "", "Path to original package")
	cfg_merge3Cmd.Flags().String("from", "", "Path to updated package")
	cfg_merge3Cmd.Flags().Bool("path-merge-key", false, "Use the path as part of the merge key when merging resources")
	cfg_merge3Cmd.Flags().String("to", "", "Path to destination package")
	cfgCmd.AddCommand(cfg_merge3Cmd)
}
