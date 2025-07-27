package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/chdman"
	"github.com/spf13/cobra"
)

var createrawCmd = &cobra.Command{
	Use:   "createraw",
	Short: "create a new compressed raw image from a raw file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createrawCmd).Standalone()

	createrawCmd.Flags().StringP("compression", "c", "", "which compression codecs to use")
	createrawCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	createrawCmd.Flags().StringP("hunksize", "hs", "", "size of each hunk, in bytes")
	createrawCmd.Flags().StringP("input", "i", "", "input file name")
	createrawCmd.Flags().StringP("inputbytes", "ib", "", "effective length of input in bytes")
	createrawCmd.Flags().StringP("inputhunks", "ih", "", "effective length of input in hunks")
	createrawCmd.Flags().StringP("inputstartbyte", "isb", "", "starting byte offset within the input")
	createrawCmd.Flags().StringP("inputstarthunk", "ish", "", "starting hunk offset within the input")
	createrawCmd.Flags().StringP("numprocessors", "np", "", "limit the number of processors to use during compression")
	createrawCmd.Flags().StringP("output", "o", "", "output file name")
	createrawCmd.Flags().StringP("outputparent", "op", "", "parent file name for output CHD")
	createrawCmd.Flags().StringP("unitsize", "us", "", "size of each unit, in bytes")
	rootCmd.AddCommand(createrawCmd)

	carapace.Gen(createrawCmd).FlagCompletion(carapace.ActionMap{
		"compression":  chdman.ActionCodecs().UniqueList(","),
		"input":        carapace.ActionFiles(),
		"output":       carapace.ActionFiles(),
		"outputparent": carapace.ActionFiles(),
	})
}
