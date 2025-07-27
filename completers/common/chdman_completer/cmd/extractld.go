package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extractldCmd = &cobra.Command{
	Use:   "extractld",
	Short: "extract laserdisc AVI from a CHD input file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extractldCmd).Standalone()

	extractldCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	extractldCmd.Flags().StringP("input", "i", "", "input file name")
	extractldCmd.Flags().StringP("inputframes", "if", "", "effective length of input in frames")
	extractldCmd.Flags().StringP("inputparent", "ip", "", "parent file name for input CHD")
	extractldCmd.Flags().StringP("inputstartframe", "isf", "", "starting frame within the input")
	extractldCmd.Flags().StringP("output", "o", "", "output file name")
	rootCmd.AddCommand(extractldCmd)

	carapace.Gen(extractldCmd).FlagCompletion(carapace.ActionMap{
		"input":       carapace.ActionFiles(),
		"inputparent": carapace.ActionFiles(),
		"output":      carapace.ActionFiles(),
	})
}
