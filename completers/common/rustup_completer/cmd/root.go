package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rustup",
	Short: "The Rust toolchain installer",
	Long:  "https://rustup.rs/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("quiet", "q", false, "Disable progress output, set log level to 'WARN' if 'RUSTUP_LOG' is unset")
	rootCmd.Flags().BoolP("verbose", "v", false, "Set log level to 'DEBUG' if 'RUSTUP_LOG' is unset")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
}
