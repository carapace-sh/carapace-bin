package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugsidedataCmd = &cobra.Command{
	Use:    "debugsidedata",
	Short:  "-c|-m|FILE REV",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugsidedataCmd).Standalone()

	debugsidedataCmd.Flags().BoolP("changelog", "c", false, "open changelog")
	debugsidedataCmd.Flags().String("dir", "", "open directory manifest")
	debugsidedataCmd.Flags().BoolP("manifest", "m", false, "open manifest")
	rootCmd.AddCommand(debugsidedataCmd)
}
