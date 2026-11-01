package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:     "store",
	Short:   "Managing the package store",
	GroupID: "store",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storeCmd).Standalone()

	storeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(storeCmd)
}
