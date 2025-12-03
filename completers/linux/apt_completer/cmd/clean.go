package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean packages and scripts from archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().Bool("dry-run", false, "perform a simulation of events taken")
	cleanCmd.Flags().BoolP("simulate", "s", false, "perform a simulation of events taken")
	rootCmd.AddCommand(cleanCmd)
}
