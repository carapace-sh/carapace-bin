package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stdCmd = &cobra.Command{
	Use:   "std",
	Short: "View standard library documentation in a browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stdCmd).Standalone()

	rootCmd.AddCommand(stdCmd)

	stdCmd.Flags().BoolP("help", "h", false, "Print help")
}
