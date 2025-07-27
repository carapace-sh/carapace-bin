package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extracthdCmd = &cobra.Command{
	Use:   "extracthd",
	Short: "extract raw hard disk file from a CHD input file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extracthdCmd).Standalone()

	extracthdCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	extracthdCmd.Flags().StringP("input", "i", "", "input file name")
	extracthdCmd.Flags().StringP("inputbytes", "ib", "", "effective length of input in bytes")
	extracthdCmd.Flags().StringP("inputhunks", "ih", "", "effective length of input in hunks")
	extracthdCmd.Flags().StringP("inputparent", "ip", "", "parent file name for input CHD")
	extracthdCmd.Flags().StringP("inputstartbyte", "isb", "", "starting byte offset within the input")
	extracthdCmd.Flags().StringP("inputstarthunk", "ish", "", "starting hunk offset within the input")
	extracthdCmd.Flags().StringP("output", "o", "", "output file name")
	rootCmd.AddCommand(extracthdCmd)

	carapace.Gen(extracthdCmd).FlagCompletion(carapace.ActionMap{
		"input":       carapace.ActionFiles(),
		"inputparent": carapace.ActionFiles(),
		"output":      carapace.ActionFiles(),
	})
}
