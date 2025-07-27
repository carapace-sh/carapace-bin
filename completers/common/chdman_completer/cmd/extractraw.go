package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extractrawCmd = &cobra.Command{
	Use:   "extractraw",
	Short: "extract raw file from a CHD input file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extractrawCmd).Standalone()

	extractrawCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	extractrawCmd.Flags().StringP("input", "i", "", "input file name")
	extractrawCmd.Flags().StringP("inputbytes", "ib", "", "effective length of input in bytes")
	extractrawCmd.Flags().StringP("inputhunks", "ih", "", "effective length of input in hunks")
	extractrawCmd.Flags().StringP("inputparent", "ip", "", "parent file name for input CHD")
	extractrawCmd.Flags().StringP("inputstartbyte", "isb", "", "starting byte offset within the input")
	extractrawCmd.Flags().StringP("inputstarthunk", "ish", "", "starting hunk offset within the input")
	extractrawCmd.Flags().StringP("output", "o", "", "output file name")
	rootCmd.AddCommand(extractrawCmd)

	carapace.Gen(extractrawCmd).FlagCompletion(carapace.ActionMap{
		"input":       carapace.ActionFiles(),
		"inputparent": carapace.ActionFiles(),
		"output":      carapace.ActionFiles(),
	})
}
