package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var _openCmd = &cobra.Command{
	Use:    "_open",
	Short:  "Open files in the workspace using any defined program",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(_openCmd).Standalone()

	_openCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	_openCmd.Flags().StringP("program-id", "p", "", "The program to use for opening")
	rootCmd.AddCommand(_openCmd)
}