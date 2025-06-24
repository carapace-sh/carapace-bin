package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Print time required to boot the machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timeCmd).Standalone()

	rootCmd.AddCommand(timeCmd)
}
