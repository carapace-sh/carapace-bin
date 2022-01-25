package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_listSettersCmd = &cobra.Command{
	Use:   "list-setters",
	Short: "[Alpha] List setters for Resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_listSettersCmd).Standalone()
	cfg_listSettersCmd.Flags().Bool("include-subst", false, "include substitutions in the output")
	cfg_listSettersCmd.Flags().Bool("markdown", false, "output as github markdown")
	cfg_listSettersCmd.Flags().BoolP("recurse-subpackages", "R", false, "list setters recursively in all the nested subpackages")
	cfgCmd.AddCommand(cfg_listSettersCmd)
}
