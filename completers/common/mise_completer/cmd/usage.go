package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var usageCmd = &cobra.Command{
	Use:    "usage",
	Short:  "Generate a usage CLI spec",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(usageCmd).Standalone()

	usageCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(usageCmd)
}
