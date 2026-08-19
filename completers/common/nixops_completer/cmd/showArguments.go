package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showArgumentsCmd = &cobra.Command{
	Use:   "show-arguments",
	Short: "print the arguments to the network expressions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showArgumentsCmd).Standalone()
	rootCmd.AddCommand(showArgumentsCmd)
}
