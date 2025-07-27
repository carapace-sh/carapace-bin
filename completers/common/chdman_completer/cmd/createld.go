package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/chdman"
	"github.com/spf13/cobra"
)

var createldCmd = &cobra.Command{
	Use:   "createld",
	Short: "create a laserdisc CHD from the input file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createldCmd).Standalone()

	createldCmd.Flags().StringP("compression", "c", "", "which compression codecs to use")
	createldCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	createldCmd.Flags().StringP("hunksize", "hs", "", "size of each hunk, in bytes")
	createldCmd.Flags().StringP("input", "i", "", "input file name")
	createldCmd.Flags().StringP("inputframes", "if", "", "effective length of input in frames")
	createldCmd.Flags().StringP("inputstartframe", "isf", "", "starting frame within the input")
	createldCmd.Flags().StringP("numprocessors", "np", "", "limit the number of processors to use during compression")
	createldCmd.Flags().StringP("output", "o", "", "output file name")
	createldCmd.Flags().StringP("outputparent", "op", "", "parent file name for output CHD")
	rootCmd.AddCommand(createldCmd)

	carapace.Gen(createldCmd).FlagCompletion(carapace.ActionMap{
		"compression":  chdman.ActionCodecs().UniqueList(","),
		"input":        carapace.ActionFiles(),
		"output":       carapace.ActionFiles(),
		"outputparent": carapace.ActionFiles(),
	})
}
