package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixos-version",
	Short: "show the NixOS version",
	Long:  "https://nixos.org/manual/nixos/stable/options.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("configuration-revision", false, "Show the configuration revision if available")
	rootCmd.Flags().Bool("hash", false, "Print the NixOS version hash (alias for --revision)")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage information")
	rootCmd.Flags().Bool("json", false, "Print a JSON representation of the versions")
	rootCmd.Flags().Bool("revision", false, "Print the NixOS version revision (alias for --hash)")
}
