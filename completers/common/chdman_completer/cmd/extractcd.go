package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extractcdCmd = &cobra.Command{
	Use:   "extractcd",
	Short: "extract CD file from a CHD input file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extractcdCmd).Standalone()

	extractcdCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	extractcdCmd.Flags().StringP("input", "i", "", "input file name")
	extractcdCmd.Flags().StringP("inputparent", "ip", "", "parent file name for input CHD")
	extractcdCmd.Flags().StringP("output", "o", "", "output file name")
	extractcdCmd.Flags().StringP("outputbin", "ob", "", "output file name for binary data")
	extractcdCmd.Flags().BoolP("splitbin", "sb", false, "output one binary file per track")
	rootCmd.AddCommand(extractcdCmd)

	carapace.Gen(extractcdCmd).FlagCompletion(carapace.ActionMap{
		"input":       carapace.ActionFiles(),
		"inputparent": carapace.ActionFiles(),
		"output":      carapace.ActionFiles(),
		"outputbin":   carapace.ActionFiles(),
	})
}
