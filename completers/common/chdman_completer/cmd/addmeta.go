package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addmetaCmd = &cobra.Command{
	Use:   "addmeta",
	Short: "add metadata to the CHD",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addmetaCmd).Standalone()

	addmetaCmd.Flags().StringP("index", "ix", "", "indexed instance of this metadata tag")
	addmetaCmd.Flags().StringP("input", "i", "", "input file name")
	addmetaCmd.Flags().BoolP("nochecksum", "nocs", false, "do not include this metadata information in the overall SHA-1")
	addmetaCmd.Flags().StringP("tag", "t", "", "4-character tag for metadata")
	addmetaCmd.Flags().StringP("valuefile", "vf", "", "file containing data to add")
	addmetaCmd.Flags().StringP("valuetext", "vt", "", "text for the metadata")
	rootCmd.AddCommand(addmetaCmd)

	carapace.Gen(addmetaCmd).FlagCompletion(carapace.ActionMap{
		"input":     carapace.ActionFiles(),
		"valuefile": carapace.ActionFiles(),
	})
}
