package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_setCmd = &cobra.Command{
	Use:   "set",
	Short: "[Alpha] Set values on Resources fields values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_setCmd).Standalone()
	cfg_setCmd.Flags().String("description", "", "annotate the field with a description of its value")
	cfg_setCmd.Flags().BoolP("recurse-subpackages", "R", false, "sets recursively in all the nested subpackages")
	cfg_setCmd.Flags().String("set-by", "", "annotate the field with who set it")
	cfg_setCmd.Flags().StringArray("values", []string{}, "optional flag, the values of the setter to be set to")
	cfg_setCmd.Flags().String("version", "", "use this version of the setter format")
	cfgCmd.AddCommand(cfg_setCmd)
}
