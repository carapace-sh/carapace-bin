package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version information for %{product}.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().Bool("gnu_format", false, "If set, write the version to stdout using the conventions described in the GNU standards.")
	versionCmd.Flags().Bool("nognu_format", false, "If set, write the version to stdout using the conventions described in the GNU standards.")
	rootCmd.AddCommand(versionCmd)
}
