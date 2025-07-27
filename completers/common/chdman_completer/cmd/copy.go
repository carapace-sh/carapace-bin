package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/chdman"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy data from one CHD to another of the same type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copyCmd).Standalone()

	copyCmd.Flags().StringP("compression", "c", "", "which compression codecs to use")
	copyCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	copyCmd.Flags().StringP("hunksize", "hs", "", "size of each hunk, in bytes")
	copyCmd.Flags().StringP("input", "i", "", "input file name")
	copyCmd.Flags().StringP("inputbytes", "ib", "", "effective length of input in bytes")
	copyCmd.Flags().StringP("inputhunks", "ih", "", "effective length of input in hunks")
	copyCmd.Flags().StringP("inputparent", "ip", "", "parent file name for input CHD")
	copyCmd.Flags().StringP("inputstartbyte", "isb", "", "starting byte offset within the input")
	copyCmd.Flags().StringP("inputstarthunk", "ish", "", "starting hunk offset within the input")
	copyCmd.Flags().StringP("numprocessors", "np", "", "limit the number of processors to use during compression")
	copyCmd.Flags().StringP("output", "o", "", "output file name")
	copyCmd.Flags().StringP("outputparent", "op", "", "parent file name for output CHD")
	rootCmd.AddCommand(copyCmd)

	carapace.Gen(copyCmd).FlagCompletion(carapace.ActionMap{
		"compression":  chdman.ActionCodecs().UniqueList(","),
		"input":        carapace.ActionFiles(),
		"inputparent":  carapace.ActionFiles(),
		"output":       carapace.ActionFiles(),
		"outputparent": carapace.ActionFiles(),
	})
}
