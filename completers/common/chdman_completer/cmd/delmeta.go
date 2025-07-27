package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var delmetaCmd = &cobra.Command{
	Use:   "delmeta",
	Short: "remove metadata from the CHD",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delmetaCmd).Standalone()

	delmetaCmd.Flags().StringP("index", "ix", "", "indexed instance of this metadata tag")
	delmetaCmd.Flags().StringP("input", "i", "", "input file name")
	delmetaCmd.Flags().StringP("tag", "t", "", "4-character tag for metadata")
	rootCmd.AddCommand(delmetaCmd)

	carapace.Gen(delmetaCmd).FlagCompletion(carapace.ActionMap{
		"input": carapace.ActionFiles(),
	})
}
