package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:     "cat",
	Short:   "Print contents of a file in a revision",
	Aliases: []string{"print"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	catCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	catCmd.Flags().StringP("revision", "r", "", "The revision to get the file contents from")
	rootCmd.AddCommand(catCmd)
}
