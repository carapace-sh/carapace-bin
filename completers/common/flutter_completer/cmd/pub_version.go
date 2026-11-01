package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print pub version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_versionCmd).Standalone()

	pub_versionCmd.Flags().Bool("color", false, "Use colors in terminal output.")
	pub_versionCmd.Flags().StringP("directory", "C", "", "Run the subcommand in the directory <dir>.")
	pub_versionCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_versionCmd.Flags().Bool("no-color", false, "Do not use colors in terminal output.")
	pub_versionCmd.Flags().BoolP("verbose", "v", false, "Print detailed logging.")
	pubCmd.AddCommand(pub_versionCmd)
}
