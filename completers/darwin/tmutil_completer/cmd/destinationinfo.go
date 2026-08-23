package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(destinationinfoCmd)
}

var destinationinfoCmd = &cobra.Command{
	Use:   "destinationinfo",
	Short: "print information about configured destinations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destinationinfoCmd).Standalone()

	destinationinfoCmd.Flags().BoolP("xml", "X", false, "Print output in XML property list format")
}