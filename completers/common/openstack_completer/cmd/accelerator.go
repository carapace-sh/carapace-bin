package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acceleratorCmd = &cobra.Command{
	Use:   "accelerator",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acceleratorCmd).Standalone()

	rootCmd.AddCommand(acceleratorCmd)
}
