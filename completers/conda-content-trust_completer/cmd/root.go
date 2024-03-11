package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "conda-content-trust",
	Short: "Signing and verification tools for Conda",
	Long:  "https://github.com/conda/conda-content-trust",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().BoolP("version", "V", false, "Show the conda-content-trust version number and exit.")
}
