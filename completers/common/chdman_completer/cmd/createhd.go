package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/chdman"
	"github.com/spf13/cobra"
)

var createhdCmd = &cobra.Command{
	Use:   "createhd",
	Short: "create a hard disk CHD from the input file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createhdCmd).Standalone()

	createhdCmd.Flags().StringP("chs", "chs", "", "specifies CHS geometry directly")
	createhdCmd.Flags().StringP("compression", "c", "", "which compression codecs to use")
	createhdCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	createhdCmd.Flags().StringP("hunksize", "hs", "", "size of each hunk, in bytes")
	createhdCmd.Flags().StringP("ident", "id", "", "name of ident file to provide CHS information")
	createhdCmd.Flags().StringP("input", "i", "", "input file name")
	createhdCmd.Flags().StringP("inputbytes", "ib", "", "effective length of input in bytes")
	createhdCmd.Flags().StringP("inputhunks", "ih", "", "effective length of input in hunks")
	createhdCmd.Flags().StringP("inputstartbyte", "isb", "", "starting byte offset within the input")
	createhdCmd.Flags().StringP("inputstarthunk", "ish", "", "starting hunk offset within the input")
	createhdCmd.Flags().StringP("numprocessors", "np", "", "limit the number of processors to use during compression")
	createhdCmd.Flags().StringP("output", "o", "", "output file name")
	createhdCmd.Flags().StringP("outputparent", "op", "", "parent file name for output CHD")
	createhdCmd.Flags().StringP("sectorsize", "ss", "", "size of each hard disk sector")
	createhdCmd.Flags().StringP("size", "s", "", "size of the output file")
	createhdCmd.Flags().StringP("template", "t", "", "use hard disk template")
	rootCmd.AddCommand(createhdCmd)

	carapace.Gen(createhdCmd).FlagCompletion(carapace.ActionMap{
		"compression":  chdman.ActionCodecs().UniqueList(","),
		"ident":        carapace.ActionFiles(),
		"input":        carapace.ActionFiles(),
		"output":       carapace.ActionFiles(),
		"outputparent": carapace.ActionFiles(),
		"template":     chdman.ActionTemplates(),
	})
}
