package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uf2Cmd = &cobra.Command{
	Use:   "uf2",
	Short: "Generate the UF2 binary with all binaries included",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uf2Cmd).Standalone()
	uf2Cmd.Flags().Bool("md5-disable", false, "Disable MD5 checksum")
	rootCmd.AddCommand(uf2Cmd)
}
