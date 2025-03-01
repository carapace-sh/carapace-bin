package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autocleanCmd = &cobra.Command{
	Use:   "autoclean",
	Short: "clean obsolete packages and scripts from archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autocleanCmd).Standalone()

	autocleanCmd.Flags().Bool("dry-run", false, "perform a simulation of events taken")
	autocleanCmd.Flags().BoolP("simulate", "s", false, "perform a simulation of events taken")
	rootCmd.AddCommand(autocleanCmd)
}
