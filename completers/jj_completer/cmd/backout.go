package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var backoutCmd = &cobra.Command{
	Use:   "backout",
	Short: "Apply the reverse of a revision on top of another revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backoutCmd).Standalone()

	backoutCmd.Flags().StringSliceP("destination", "d", []string{}, "The revision to apply the reverse changes on top of")
	backoutCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	backoutCmd.Flags().StringP("revision", "r", "", "The revision to apply the reverse of")
	rootCmd.AddCommand(backoutCmd)
}
