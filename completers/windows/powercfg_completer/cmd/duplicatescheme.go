package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duplicateschemeCmd = &cobra.Command{
	Use:   "duplicatescheme",
	Short: "duplicate a power plan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duplicateschemeCmd).Standalone()
	rootCmd.AddCommand(duplicateschemeCmd)
}
