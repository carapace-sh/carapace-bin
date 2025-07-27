package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/chdman"
	"github.com/spf13/cobra"
)

var createdvdCmd = &cobra.Command{
	Use:   "createdvd",
	Short: "create a DVD CHD from the input file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createdvdCmd).Standalone()

	createdvdCmd.Flags().StringP("compression", "c", "", "which compression codecs to use")
	createdvdCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	createdvdCmd.Flags().StringP("hunksize", "hs", "", "size of each hunk, in bytes")
	createdvdCmd.Flags().StringP("input", "i", "", "input file name")
	createdvdCmd.Flags().StringP("inputbytes", "ib", "", "effective length of input in bytes")
	createdvdCmd.Flags().StringP("inputhunks", "ih", "", "effective length of input in hunks")
	createdvdCmd.Flags().StringP("inputstartbyte", "isb", "", "starting byte offset within the input")
	createdvdCmd.Flags().StringP("inputstarthunk", "ish", "", "starting hunk offset within the input")
	createdvdCmd.Flags().StringP("numprocessors", "np", "", "limit the number of processors to use during compression")
	createdvdCmd.Flags().StringP("output", "o", "", "output file name")
	createdvdCmd.Flags().StringP("outputparent", "op", "", "parent file name for output CHD")
	rootCmd.AddCommand(createdvdCmd)

	carapace.Gen(createdvdCmd).FlagCompletion(carapace.ActionMap{
		"compression":  chdman.ActionCodecs().UniqueList(","),
		"input":        carapace.ActionFiles(),
		"output":       carapace.ActionFiles(),
		"outputparent": carapace.ActionFiles(),
	})
}
