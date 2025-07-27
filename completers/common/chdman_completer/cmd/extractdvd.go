package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extractdvdCmd = &cobra.Command{
	Use:   "extractdvd",
	Short: "extract DVD file from a CHD input file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extractdvdCmd).Standalone()

	extractdvdCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	extractdvdCmd.Flags().StringP("input", "i", "", "input file name")
	extractdvdCmd.Flags().StringP("inputbytes", "ib", "", "effective length of input in bytes")
	extractdvdCmd.Flags().StringP("inputhunks", "ih", "", "effective length of input in hunks")
	extractdvdCmd.Flags().StringP("inputparent", "ip", "", "parent file name for input CHD")
	extractdvdCmd.Flags().StringP("inputstartbyte", "isb", "", "starting byte offset within the input")
	extractdvdCmd.Flags().StringP("inputstarthunk", "ish", "", "starting hunk offset within the input")
	extractdvdCmd.Flags().StringP("output", "o", "", "output file name")
	rootCmd.AddCommand(extractdvdCmd)

	carapace.Gen(extractdvdCmd).FlagCompletion(carapace.ActionMap{
		"input":       carapace.ActionFiles(),
		"inputparent": carapace.ActionFiles(),
		"output":      carapace.ActionFiles(),
	})
}
