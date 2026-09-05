package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manage Node.js versions. Deprecated in favour of `pnpm runtime`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(envCmd).Standalone()

	envCmd.Flags().BoolP("global", "g", false, "Manage Node.js versions globally")
	envCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	envCmd.Flags().Bool("remote", false, "Accepted for surface parity with pnpm, which declares the option but never reads it")
	envCmd.Flag("remote").Hidden = true
	rootCmd.AddCommand(envCmd)
}
