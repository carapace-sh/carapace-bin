package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enumdependCmd = &cobra.Command{
	Use:   "enumdepend",
	Short: "enumerate dependent services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enumdependCmd).Standalone()
	rootCmd.AddCommand(enumdependCmd)
}
