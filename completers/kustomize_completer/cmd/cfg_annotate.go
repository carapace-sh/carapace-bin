package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_annotateCmd = &cobra.Command{
	Use:   "annotate",
	Short: "[Alpha] Set an annotation on Resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_annotateCmd).Standalone()
	cfg_annotateCmd.Flags().String("apiVersion", "", "Resource apiVersion to annotate")
	cfg_annotateCmd.Flags().String("kind", "", "Resource kind to annotate")
	cfg_annotateCmd.Flags().StringSlice("kv", []string{}, "annotation as KEY=VALUE")
	cfg_annotateCmd.Flags().String("name", "", "Resource name to annotate")
	cfg_annotateCmd.Flags().String("namespace", "", "Resource namespace to annotate")
	cfg_annotateCmd.Flags().BoolP("recurse-subpackages", "R", false, "add annotations recursively in all the nested subpackages")
	cfgCmd.AddCommand(cfg_annotateCmd)
}
