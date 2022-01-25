package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfg_fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "[Alpha] Format yaml configuration files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfg_fmtCmd).Standalone()
	cfg_fmtCmd.Flags().Bool("keep-annotations", false, "if true, keep index and filename annotations set on Resources.")
	cfg_fmtCmd.Flags().Bool("override", false, "if true, override existing filepath annotations.")
	cfg_fmtCmd.Flags().String("pattern", "%n_%k.yaml", "pattern to use for generating filenames for resources -- may contain the following")
	cfg_fmtCmd.Flags().BoolP("recurse-subpackages", "R", false, "formats resource files recursively in all the nested subpackages")
	cfg_fmtCmd.Flags().Bool("set-filenames", false, "if true, set default filenames on Resources without them")
	cfg_fmtCmd.Flags().Bool("use-schema", false, "if true, uses openapi resource schema to format resources.")
	cfgCmd.AddCommand(cfg_fmtCmd)
}
