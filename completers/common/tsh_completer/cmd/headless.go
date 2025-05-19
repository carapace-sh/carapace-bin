package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var headlessCmd = &cobra.Command{
	Use:   "headless",
	Short: "Headless authentication commands.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(headlessCmd).Standalone()

	rootCmd.AddCommand(headlessCmd)
}
