package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var allowCmd = &cobra.Command{
	Use:   "allow",
	Short: "add allow rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(allowCmd)

	rootCmd.AddCommand(allowCmd)
}
