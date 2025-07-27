package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/chdman"
	"github.com/spf13/cobra"
)

var createcdCmd = &cobra.Command{
	Use:   "createcd",
	Short: "create a new compressed CD image from a raw file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createcdCmd).Standalone()

	createcdCmd.Flags().StringP("compression", "c", "", "which compression codecs to use")
	createcdCmd.Flags().StringP("hunksize", "hs", "", "size of each hunk, in bytes")
	createcdCmd.Flags().StringP("input", "i", "", "input file name")
	createcdCmd.Flags().StringP("numprocessors", "np", "", "limit the number of processors to use during compression")
	createcdCmd.Flags().StringP("output", "o", "", "output file name")
	createcdCmd.Flags().StringP("outputparent", "op", "", "parent file name for output CHD")
	rootCmd.AddCommand(createcdCmd)

	carapace.Gen(createcdCmd).FlagCompletion(carapace.ActionMap{
		"compression":  chdman.ActionCodecs().UniqueList(","),
		"input":        carapace.ActionFiles(),
		"output":       carapace.ActionFiles(),
		"outputparent": carapace.ActionFiles(),
	})
}
