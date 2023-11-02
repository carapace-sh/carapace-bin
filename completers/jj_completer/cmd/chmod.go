package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var chmodCmd = &cobra.Command{
	Use:   "chmod",
	Short: "Sets or removes the executable bit for paths in the repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chmodCmd).Standalone()

	chmodCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	chmodCmd.Flags().StringP("revision", "r", "", "The revision to update")
	rootCmd.AddCommand(chmodCmd)
}
