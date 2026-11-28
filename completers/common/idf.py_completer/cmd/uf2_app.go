package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uf2AppCmd = &cobra.Command{
	Use:   "uf2-app",
	Short: "Generate a UF2 binary for the application only",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uf2AppCmd).Standalone()

	uf2AppCmd.Flags().Bool("md5-disable", false, "Disable MD5 checksum")
	rootCmd.AddCommand(uf2AppCmd)
}
