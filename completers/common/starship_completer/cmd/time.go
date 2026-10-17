package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:    "time",
	Short:  "Prints time in milliseconds",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timeCmd).Standalone()

	timeCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(timeCmd)
}
