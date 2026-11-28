package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dfuCmd = &cobra.Command{
	Use:   "dfu",
	Short: "Build the DFU binary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dfuCmd).Standalone()
	dfuCmd.Flags().String("part-size", "", "Overwrite default partition size of mkdfu.py")
	rootCmd.AddCommand(dfuCmd)
}
