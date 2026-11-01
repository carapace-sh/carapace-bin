package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up pnpm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupCmd).Standalone()

	setupCmd.Flags().BoolP("force", "f", false, "Override the PNPM_HOME env variable in case it already exists")
	setupCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(setupCmd)
}
