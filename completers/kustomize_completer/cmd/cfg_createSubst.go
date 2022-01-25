package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_createSubstCmd = &cobra.Command{
	Use:   "create-subst",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_createSubstCmd).Standalone()
	cfg_createSubstCmd.Flags().String("field", "", "name of the field to set -- e.g. --field image")
	cfg_createSubstCmd.Flags().String("field-value", "", "value of the field to create substitution for -- e.g. --field-value nginx:0.1.0")
	cfg_createSubstCmd.Flags().String("pattern", "", "substitution pattern")
	cfg_createSubstCmd.Flags().BoolP("recurse-subpackages", "R", false, "creates substitution recursively in all the nested subpackages")
	cfgCmd.AddCommand(cfg_createSubstCmd)
}
