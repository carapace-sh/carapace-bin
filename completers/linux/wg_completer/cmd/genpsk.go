package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var genpskCmd = &cobra.Command{
	Use:   "genpsk",
	Short: "Generates a new preshared key and writes it to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genpskCmd).Standalone()
	rootCmd.AddCommand(genpskCmd)
}