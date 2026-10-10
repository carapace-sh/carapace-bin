package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugdataCmd = &cobra.Command{
	Use:    "debugdata",
	Short:  "-c|-m|FILE REV",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugdataCmd).Standalone()

	debugdataCmd.Flags().BoolP("changelog", "c", false, "open changelog")
	debugdataCmd.Flags().String("dir", "", "open directory manifest")
	debugdataCmd.Flags().BoolP("manifest", "m", false, "open manifest")
	rootCmd.AddCommand(debugdataCmd)
}
