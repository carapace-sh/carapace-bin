package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Print a JWT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(jwtCmd).Standalone()

	rootCmd.AddCommand(jwtCmd)
}
