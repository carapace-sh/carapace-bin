package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "maturin",
	Short: "Build and publish crates with pyo3, cffi and uniffi bindings as well as rust binaries as python packages",
	Long:  "https://www.maturin.rs/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")
	rootCmd.Flags().CountP("verbose", "v", "Use verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
}
