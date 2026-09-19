package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugnodemapCmd = &cobra.Command{
	Use:    "debugnodemap",
	Short:  "-c|-m|FILE",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugnodemapCmd).Standalone()

	debugnodemapCmd.Flags().BoolP("changelog", "c", false, "open changelog")
	debugnodemapCmd.Flags().Bool("check", false, "check that the data on disk data are correct.")
	debugnodemapCmd.Flags().String("dir", "", "open directory manifest")
	debugnodemapCmd.Flags().Bool("dump-disk", false, "dump on-disk data on stdout")
	debugnodemapCmd.Flags().Bool("dump-new", false, "write a (new) persistent binary nodemap on stdout")
	debugnodemapCmd.Flags().BoolP("manifest", "m", false, "open manifest")
	debugnodemapCmd.Flags().Bool("metadata", false, "display the on disk meta data for the nodemap")
	rootCmd.AddCommand(debugnodemapCmd)
}
