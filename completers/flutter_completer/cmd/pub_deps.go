package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Print package dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_depsCmd).Standalone()

	pub_depsCmd.Flags().StringP("device-id", "d", "", "Target device id or name (prefixes allowed).")
	pub_depsCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_depsCmd.Flags().Bool("suppress-analytics", false, "Suppress analytics reporting when this command runs.")
	pub_depsCmd.Flags().BoolP("verbose", "v", false, "Noisy logging, including all shell commands executed.")
	pub_depsCmd.Flags().Bool("version", false, "Reports the version of this tool.")
	pubCmd.AddCommand(pub_depsCmd)
}
