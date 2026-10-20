package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustCmd = &cobra.Command{
	Use:   "trust",
	Short: "Mark a config file as trusted",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustCmd).Standalone()

	trustCmd.Flags().BoolP("all", "a", false, "Trust all config files in the current directory, its parents, and its subdirectories")
	trustCmd.Flags().BoolP("help", "h", false, "Print help")
	trustCmd.Flags().Bool("ignore", false, "Do not trust this config and ignore it in the future")
	trustCmd.Flags().Bool("show", false, "Show the trust status of config files in the current directory and its parents")
	trustCmd.Flags().Bool("untrust", false, "Remove explicit trust for this config")
	rootCmd.AddCommand(trustCmd)

	carapace.Gen(trustCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
