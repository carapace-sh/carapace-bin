package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ociCmd = &cobra.Command{
	Use:   "oci",
	Short: "[experimental] Build, push, and run OCI images of a project's tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ociCmd).Standalone()

	ociCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(ociCmd)
}
