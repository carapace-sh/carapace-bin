package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printConfigCmd = &cobra.Command{
	Use:   "print-config",
	Short: "Prints the computed starship configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printConfigCmd).Standalone()

	printConfigCmd.Flags().BoolP("default", "d", false, "Print the default instead of the computed config")
	printConfigCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(printConfigCmd)
}
