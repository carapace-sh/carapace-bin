package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpmetaCmd = &cobra.Command{
	Use:   "dumpmeta",
	Short: "dump metadata from the CHD to stdout or to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpmetaCmd).Standalone()

	dumpmetaCmd.Flags().BoolP("force", "f", false, "force overwriting an existing file")
	dumpmetaCmd.Flags().StringP("index", "ix", "", "indexed instance of this metadata tag")
	dumpmetaCmd.Flags().StringP("input", "i", "", "input file name")
	dumpmetaCmd.Flags().StringP("output", "o", "", "output file name")
	dumpmetaCmd.Flags().StringP("tag", "t", "", "4-character tag for metadata")
	rootCmd.AddCommand(dumpmetaCmd)

	carapace.Gen(dumpmetaCmd).FlagCompletion(carapace.ActionMap{
		"input":  carapace.ActionFiles(),
		"output": carapace.ActionFiles(),
	})
}
